package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

var dataFolder string

type Task struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

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

func getTaskFile(id string) string {
	return filepath.Join(dataFolder, fmt.Sprintf("%s.json", id))
}

func writeToFile(filepath string, content []byte) error {
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("Unable to create file : %v", err)
	}
	defer file.Close()

	file.Write(content)
	return nil
}

func (t *Task) Write() error {
	jsonBytes, err := json.Marshal(t)
	if err != nil {
		return fmt.Errorf("Unable to write: %v", err)
	}

	writeToFile(getTaskFile(t.ID), jsonBytes)
	return nil
}

func GetTask(taskId string) (Task, error) {
	var task Task

	fileContent, err := os.ReadFile(getTaskFile(taskId))
	if err != nil {
		return task, fmt.Errorf("Unable to read file: %v", err)
	}

	if err = json.Unmarshal(fileContent, &task); err != nil {
		return task, fmt.Errorf("Unable to unmarshal data: %v", err)
	}
	return task, nil
}
