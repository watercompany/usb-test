package utils

import (
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

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

	log.Printf("File created: %s \n", fileName)

	return fileName, nil
}

// GenerateRandomString is a helper func for generating
// random string of the given input length
// returns the generated string
func GenerateRandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	str := b.String()

	return str
}