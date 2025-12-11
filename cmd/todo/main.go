package main

import (
	"fmt"
	"os"
	"todo/internal/httpapi"
	"todo/internal/tasks"
)

func main() {

	connString := "postgres://postgres:1234@localhost:5432/todo?sslmode=disable"

	_, err := tasks.InitDB(connString)
	if err != nil {
		panic(err)
	}

	repo := tasks.NewSQLRepository(tasks.DB)

	args := os.Args[1:]

	// --- CLI MODE ---
	if len(args) > 0 {
		runCLI(repo, args)
		return
	}

	// --- SERVER MODE ---
	fmt.Println("Server HTTP attivo su http://localhost:8080")
	httpapi.Startserver(repo)
}
