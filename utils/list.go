package utils

import (
	"log"
	"os"
	"path/filepath"
)

var dataFolder string

func init() {
	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("Unable to access users directory.")
	}

	dataFolder = filepath.Join(userHomeDir, ".tasks")
	if _, err := os.Stat(dataFolder); os.IsNotExist(err) {
		err := os.Mkdir(dataFolder, 0700)
		if err != nil {
			log.Fatal("Unable to create data folder.")
		}
	}
}
