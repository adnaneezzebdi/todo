package tasks

import (
	"errors"
	"time"
)

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

func Complete(id int) error {
	tasks, err := LoadTasks()
	if err != nil {
		return err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			now := time.Now()
			tasks[i].Done = true
			tasks[i].CompleteAt = &now
			return SaveTasks(tasks)
		}
	}
	return errors.New("task non trovata")
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
		return errors.New("task non trovata")
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
}
