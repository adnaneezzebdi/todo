package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if err := loadTasks(); err != nil {
		println("errore caricamento:", err.Error())
		return
	}

	if len(os.Args) < 2 {
		println("comandi disponibili: add, list, complete,delete")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		handleAdd(os.Args[2:])
	case "list":
		handleList()
	case "complete":
		handleComplete(os.Args[2:])
	case "delete":
		handleDelete(os.Args[2:])
	default:
		println("comando sbagliato")
	}
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
	data, err := json.MarshalIndent(tasks, "", "  ") //converti lo slice in json
	if err != nil {
		return err
	}

	err = os.WriteFile("tasks.json", data, 0644) //scrivi nel file
	if err != nil {
		return err
	}

	return nil
}

func handleAdd(args []string) {
	if len(args) == 0 {
		println("Usage: todo add <titolo>")
		return
	}

	title := strings.Join(args, " ")

	id := len(tasks) + 1

	newTask := Task{
		ID:        id,
		Title:     title,
		DueDate:   time.Now(),
		Completed: false,
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(); err != nil {
		println("Errore nel salvataggio:", err.Error())
		return
	}

	println("Task aggiunta:", title)
}

func handleList() {
	if len(tasks) == 0 {
		println("Nessuna task trovata.")
		return
	}

	for _, task := range tasks {
		status := "[ ]"
		if task.Completed {
			status = "[x]"
		}

		println(fmt.Sprintf("%d. %s %s", task.ID, status, task.Title))
	}
}

func handleComplete(args []string) {
	if len(args) == 0 {
		println("Usage: todo complete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		println("L'ID deve essere un numero.")
		return
	}

	found := false
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Completed = true
			found = true
			break
		}
	}

	if !found {
		println("Task non trovata con ID:", id)
		return
	}

	if err := saveTasks(); err != nil {
		println("Errore nel salvataggio:", err.Error())
		return
	}

	println("Task completata!")
}

func handleDelete(args []string) {
	if len(args) == 0 {
		println("usage: todo delete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		println("l'ID deve essere un numero")
		return
	}

	index := -1

	for i := range tasks {
		if tasks[i].ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		println("task non trovata con ID:", id)
		return
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	if err := saveTasks(); err != nil {
		println("errore nel salvataggio:", err.Error())
		return
	}

	println("task eliminata")
}
