package main

import (
	"fmt"
	"todo/internal/tasks"
)

func runCLI(repo *tasks.SQLRepository, args []string) {
	if len(args) < 1 {
		fmt.Println("Comandi disponibili: list, add, done, delete")
		return
	}

	switch args[0] {

	case "list":
		tasksList, err := repo.List()
		if err != nil {
			fmt.Println("Errore:", err)
			return
		}

		for _, t := range tasksList {
			status := " "
			if t.Done {
				status = "✓"
			}
			fmt.Printf("[%d] %s %s\n", t.ID, status, t.Title)
		}

	case "add":
		if len(args) < 2 {
			fmt.Println("Uso: todo add \"titolo\"")
			return
		}

		title := args[1]
		id, err := repo.Add(title)
		if err != nil {
			fmt.Println("Errore:", err)
			return
		}

		fmt.Println("Task aggiunta con ID:", id)

	case "done":
		if len(args) < 2 {
			fmt.Println("Uso: todo done <id>")
			return
		}

		var id int
		fmt.Sscanf(args[1], "%d", &id)

		if err := repo.Complete(id); err != nil {
			fmt.Println("Errore:", err)
			return
		}

		fmt.Println("Task completata!")

	case "delete":
		if len(args) < 2 {
			fmt.Println("Uso: todo delete <id>")
			return
		}

		var id int
		fmt.Sscanf(args[1], "%d", &id)

		if err := repo.Delete(id); err != nil {
			fmt.Println("Errore:", err)
			return
		}

		fmt.Println("Task eliminata!")

	default:
		fmt.Println("Comando sconosciuto:", args[0])
	}
}
