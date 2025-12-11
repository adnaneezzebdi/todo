package tasks

/*import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

const dataDir = "data"
const tasksFile = "tasks.json"

func getFilePath() string {
	return filepath.Join(dataDir, tasksFile)
}

func LoadTasks() ([]Task, error) {
	path := getFilePath()

	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func SaveTasks(tasks []Task) error {
	path := getFilePath()

	if err := os.MkdirAll(dataDir, 0755); err != nil {
		return err
	}

	bytes, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(path, bytes, 0644)
}*/
