package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	ID        int
	Title     string
	DueDate   time.Time
	Completed bool
}

var tasks []Task

func main() {

}

func loadTasks() error { // controlla se il file esiste
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) { //creiamo il file siccome non esiste
		empty := []byte("[]")
		err := os.WriteFile("tasks.json", empty, 0644)
		if err != nil {
			return err
		}
		tasks = []Task{} //slice vuoto in memoria
		return nil
	}

	if err != nil { //vuol dire che è un altro tipo di errore, diverso da Stat
		return err
	}

	data, err := os.ReadFile("tasks.json") //legge il contenuto del file
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &tasks) //converte il json dentro lo slice
	if err != nil {
		return err
	}
	return nil
}

func saveTasks() error {
	return nil
}
