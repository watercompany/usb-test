package cmd

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"time"

	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
	"github.com/watercompany/usb-test/utils"
)

type PathError struct {
	Path  string `json:"path"`
	Type  string `json:"type"`
	Error error  `json:"error"`
}

const (
	MB = 1024 * 1024
)

var (
	byteSize       = 1024 * 1024
	shaFileName    = fmt.Sprintf("%d-SHA256", byteSize/1024)
	mediaDirectory = "/mnt/"
	testErrors     = []PathError{}
)

func RunTest(ctx *cli.Context, numSimReadWrite int) error {
	// lsblkJSON, err := ParseLsblk()
	// if err != nil {
	// 	return err
	// }

	// var mountPoints []string
	// var shaFiles [][]byte

	// for _, device := range lsblkJSON.BlockDevices {
	// 	log.Printf("%+v\n", device)
	// 	// check if it is a usb device and if so get the mount points
	// 	for _, child := range device.Children {
	// 		mountPoints = append(mountPoints, child.Name)
	// 		token := make([]byte, byteSize)
	// 		rand.Read(token)
	// 		shaFiles = append(shaFiles, token)
	// 	}
	// }

	var shaFiles [][]byte
	mountPoints, err := utils.ListDirectories(mediaDirectory)
	if err != nil {
		return err
	}

	n := len(mountPoints)

	for i := 0; i < n; i++ {
		mountPoints[i] = filepath.Join(mountPoints[i], shaFileName)
		token := make([]byte, byteSize)
		rand.Read(token)
		shaFiles = append(shaFiles, token)
	}

	// write to files
	log.Println("-------------------STAGE 1---------------------")
	// log.Println("Creating files: ...")
	writeDuration := writeToMounts(shaFiles, mountPoints, numSimReadWrite)

	// log.Println("Files created.")
	log.Printf("Time taken to write: %s\n", writeDuration)
	writeSpeed := float64(1000000000*byteSize) / float64(writeDuration.Nanoseconds()*MB)
	log.Printf("Write speed: %f MB/s\n", writeSpeed)
	log.Println("-----------------------------------------------")

	// read from files
	log.Println("-------------------STAGE 2---------------------")
	readDuration := readFromMounts(shaFiles, mountPoints, numSimReadWrite)
	readSpeed := float64(1000000000.0*byteSize) / float64(readDuration.Nanoseconds()*MB)
	log.Printf("Time taken to read: %s\n", readDuration)
	log.Printf("Read speed: %f MB/s\n", readSpeed)
	log.Println("-----------------------------------------------")

	log.Println("-------------------Summary---------------------")
	log.Printf("Number of files: %d", n)
	log.Printf("Size of each file: %f MB", float64(byteSize)/MB)
	log.Println("-----------------------------------------------")

	// clean up files
	for _, mountPoint := range mountPoints {
		err = deleteFile(mountPoint)
		if err != nil {
			// log.Println("Unable to delete: ", mountPoint)
			testErrors = append(testErrors, PathError{Path: mountPoint, Error: err, Type: "delete"})
		}
	}

	if len(testErrors) > 0 {
		log.Println("-------------------Errors---------------------")
		for _, testError := range testErrors {
			fmt.Printf("Error: %s, %s, %+v", testError.Path, testError.Type, testError.Error)
		}
		log.Println("-----------------------------------------------")
	}

	return nil
}

func deleteFile(path string) error {
	err := os.Remove(path) // remove a single file
	if err != nil {
		return err
	}
	return nil
}

func writeToMounts(shaFiles [][]byte, mountPoints []string, numWorkers int) *time.Duration {
	// write file to path
	start := time.Now()

	numberOfFileJobs := len(mountPoints)
	jobs := make(chan int, numberOfFileJobs)
	results := make(chan int, numberOfFileJobs)

	for w := 1; w <= numWorkers; w++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for j := range jobs {
				// time.Sleep(time.Second)
				// fmt.Println("worker", id, "finished job", j)
				shaFile := shaFiles[j]
				writePath := mountPoints[j]
				// fmt.Println(writePath)

				createdFilePath, err := utils.CreateFile(writePath, true)
				if err != nil {
					testErrors = append(testErrors, PathError{Path: writePath, Error: err, Type: "write"})
					continue
				}
				file, err := os.OpenFile(createdFilePath, os.O_RDWR, 0644)
				if err != nil {
					testErrors = append(testErrors, PathError{Path: writePath, Error: err, Type: "write"})
					continue
				}

				// write shaFile to file
				_, err = file.Write(shaFile)
				if err != nil {
					testErrors = append(testErrors, PathError{Path: writePath, Error: err, Type: "write"})
					continue
				}
				file.Close()
				results <- j
			}
		}(w, jobs, results)
	}

	for j := 0; j < numberOfFileJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// get results
	for a := 1; a <= numberOfFileJobs; a++ {
		<-results
	}

	duration := time.Since(start)
	return &duration
}

func readFromMounts(shaFiles [][]byte, mountPoints []string, numWorkers int) *time.Duration {
	// write file to path
	start := time.Now()

	numberOfFileJobs := len(mountPoints)
	jobs := make(chan int, numberOfFileJobs)
	results := make(chan int, numberOfFileJobs)

	for w := 1; w <= numWorkers; w++ {
		go func(id int, jobs <-chan int, results chan<- int) {
			for j := range jobs {
				shaFile := shaFiles[j]
				readPath := mountPoints[j]

				file, err := os.OpenFile(readPath, os.O_RDWR, 0644)
				if err != nil {
					testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
					continue
				}

				token := make([]byte, byteSize)
				readByteLength, err := file.Read(token)
				if err != nil {
					testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
					continue
				}

				if readByteLength != byteSize {
					err := errors.Errorf("length of file and token not equal %d, %d\n", readByteLength, byteSize)
					testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
				}

				if !reflect.DeepEqual(token, shaFile) {
					err := errors.Errorf("file has a different content\n %x, \n%x\n", token, shaFile)
					testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
				}
				file.Close()
				// fmt.Println("worker", id, "finished job", j)
				results <- j
			}
		}(w, jobs, results)
	}

	for j := 0; j < numberOfFileJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// get results
	for a := 1; a <= numberOfFileJobs; a++ {
		<-results
	}

	// for i, shaFile := range shaFiles {
	// 	readPath := filepath.Join(mountPoints[i], shaFileName)

	// 	file, err := os.OpenFile(readPath, os.O_RDWR, 0644)
	// 	if err != nil {
	// 		testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
	// 		continue
	// 	}
	// 	defer file.Close()

	// 	token := make([]byte, byteSize)
	// 	readByteLength, err := file.Read(token)
	// 	if err != nil {
	// 		testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
	// 		continue
	// 	}

	// 	if readByteLength != byteSize {
	// 		err := errors.Errorf("length of file and token not equal %d, %d\n", readByteLength, byteSize)
	// 		testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
	// 	}

	// 	if !reflect.DeepEqual(token, shaFile) {
	// 		err := errors.Errorf("file has a different content\n %x, \n%x\n", token, shaFile)
	// 		testErrors = append(testErrors, PathError{Path: readPath, Error: err, Type: "read"})
	// 	}

	// }

	duration := time.Since(start)
	return &duration
}

// {
// 	"blockdevices": [
// 	   {"name":"vda", "maj:min":"254:0", "rm":false, "size":"59.6G", "ro":false, "type":"disk", "mountpoint":null,
// 		  "children": [
// 			 {"name":"vda1", "maj:min":"254:1", "rm":false, "size":"59.6G", "ro":false, "type":"part", "mountpoint":"/etc/hosts"}
// 		  ]
// 	   }
// 	]
//  }

// LSBLK is a wrapper for the lsblk output
type LSBLK struct {
	BlockDevices []Device `json:"blockdevices"`
}

// Device is a wrapper for device in LSBLK
type Device struct {
	Children []Child `json:"children"`
}

// Child is a wrapper for children in Device
type Child struct {
	Name       string  `json:"name"`
	MajMin     string  `json:"maj:min"`
	Rm         bool    `json:"rm"`
	Size       float64 `json:"size"`
	Ro         bool    `json:"ro"`
	Type       string  `json:"type"`
	MountPoint string  `json:"mountpoint"`
}

// ParseLsblk returns json format of `lsblk --json` command
func ParseLsblk() (LSBLK, error) {
	var parsedLSBLK LSBLK
	lsblkOutput, err := utils.RunCMD("lsblk", "--json")
	if err != nil {
		return LSBLK{}, err
	}
	json.Unmarshal([]byte(lsblkOutput), &parsedLSBLK)
	return parsedLSBLK, nil
}
