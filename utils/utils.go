package utils

import (
	"crypto/sha256"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/pkg/errors"
)

// CreateFolder creates the folder in the specified `path`
// Print success info log on successfully ran command, return error if fail
func CreateFolder(path string, force bool) error {
	if force {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Println("Error creating directory")
			return errors.Errorf("Failed to create directory: %v \n", err)
		}
		log.Printf("Directory created on %s \n", path)
	} else {
		if _, err := os.Stat(path); os.IsNotExist(err) {
			err := os.MkdirAll(path, 0755)
			if err != nil {
				log.Println("Error creating directory")
				return errors.Errorf("Failed to create directory: %v \n", err)
			}
			log.Printf("Directory created on %s \n", path)
		} else {
			log.Printf("Directory already exists \n")
			return errors.Errorf("Directory already exists \n")
		}
	}

	return nil
}

// CreateFile creates a file
// Print success info log on successfully ran command, return error if fail
func CreateFile(fileName string, force bool) (string, error) {

	if force {
		var file, err = os.Create(fileName)
		if err != nil {
			log.Println("Error creating file")
			return "", errors.Errorf("Failed to create file: %v \n", err)
		}
		defer file.Close()
	} else {
		// check if file exists
		var _, err = os.Stat(fileName)

		// create file if not exists
		if os.IsNotExist(err) {
			var file, err = os.Create(fileName)
			if err != nil {
				log.Println("Error creating file")
				return "", errors.Errorf("Failed to create file: %v \n", err)
			}
			defer file.Close()
		} else {
			log.Printf("File already exists \n")
			return "", errors.Errorf("File already exists \n")
		}
	}

	return fileName, nil
}

// RunCMD runs shell command and returns output and error
func RunCMD(command string, args ...string) (string, error) {
	cmd := exec.Command(command, args...)

	stdout, err := cmd.Output()

	if err != nil {
		return "", err
	}
	return string(stdout), nil
}

// NewSHA256 generates sha256 hash
func NewSHA256(data []byte) []byte {
	hash := sha256.Sum256(data)
	return hash[:]
}

// ListDirectories returns slice of directories in a specific path
func ListDirectories(path string) ([]string, error) {
	var dirList []string
	files, err := ioutil.ReadDir(path)

	if err != nil {

		return dirList, err
	}

	for _, f := range files {

		if f.IsDir() {
			dirList = append(dirList, filepath.Join(path, f.Name()))
		}
	}
	return dirList, nil
}
