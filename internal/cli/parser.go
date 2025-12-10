package cli

import (
	"fmt"
	"os"
	"strconv"
	"todo/internal/tasks"
)

func Run() {
	if len(os.Args) < 2 {
		printHelp()
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
		println("comando sconosciuto:", command)
		printHelp()
	}
}

func handleAdd(args []string) {
	if len(args) == 0 {
		println("usage: todo add <titolo>")
		return
	}

	title := args[0] // v1: solo una parola
	if err := tasks.Add(title); err != nil {
		fmt.Println("errore:", err)
		return
	}

	println("task aggiunta!")
}

func handleList() {
	ts, err := tasks.List()
	if err != nil {
		fmt.Println("errore:", err)
		return
	}

	if len(ts) == 0 {
		println("Nessuna task trovata.")
		return
	}

	for _, t := range ts {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s %d: %s\n", status, t.ID, t.Title)
	}
}

func handleComplete(args []string) {
	if len(args) == 0 {
		println("usage: todo complete <id>")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		println("l'ID deve essere un numero")
		return
	}

	if err := tasks.Complete(id); err != nil {
		fmt.Println("errore:", err)
		return
	}

	println("task completata!")
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

	if err := tasks.Delete(id); err != nil {
		fmt.Println("errore:", err)
		return
	}

	println("task eliminata!")
}

func printHelp() {
	println("Comandi disponibili:")
	println("  todo add <titolo>")
	println("  todo list")
	println("  todo complete <id>")
	println("  todo delete <id>")
}
