package tasks

/*import (
	"errors"
	"time"
)

var ErrNotFound = errors.New("task not found")

func Add(title string) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	newTask := Task{
		ID:        nextID(tasks),
		Title:     title,
		Done:      false,
		CreatedAt: time.Now(),
	}
	tasks = append(tasks, newTask)
	return SaveTasks(tasks)
}

func List() ([]Task, error) {
	return LoadTasks()
}

func Complete(id int) (Task, error) {
	tasks, err := LoadTasks()
	if err != nil {
		return Task{}, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			now := time.Now()
			tasks[i].Done = true
			tasks[i].CompletedAt = &now

			err = SaveTasks(tasks)
			if err != nil {
				return Task{}, err
			}

			return tasks[i], nil
		}
	}

	return Task{}, ErrNotFound
}

func Delete(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	index := -1

	for i := range tasks {
		if tasks[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		return ErrNotFound
	}

	tasks = append(tasks[:index], tasks[index+1:]...)
	return SaveTasks(tasks)
}

func nextID(tasks []Task) int {
	max := 0
	for _, t := range tasks {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
} */
