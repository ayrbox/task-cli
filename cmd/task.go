package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
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

func getFiles() ([]string, error) {
	files, err := os.ReadDir(dataFolder)
	if err != nil {
		return nil, fmt.Errorf("Unable to read data folders: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		fileName := file.Name()

		// remove file extension
		fileId, _ := strings.CutSuffix(fileName, filepath.Ext(fileName))

		fileNames = append(fileNames, fileId)
	}

	return fileNames, nil
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

func GetTasks() ([]Task, error) {
	var tasks []Task

	files, err := getFiles()
	if err != nil {
		return tasks, fmt.Errorf("Unable read files: %v", err)
	}

	for _, filename := range files {
		t, err := GetTask(filename)
		if err != nil {
			fmt.Printf("Unable to read task: %s\n", filename)
			continue
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
