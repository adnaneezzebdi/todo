package main

import (
	"fmt"
	"os"
	"todo/internal/httpapi"
	"todo/internal/tasks"
)

func main() {

	connString := "postgresql://dbtodo_wiwr_user:d55Uobzael8R79BDlNwfeFlDBmN7PvSW@dpg-d4vt6cruibrs73dn44v0-a.frankfurt-postgres.render.com/dbtodo_wiwr?sslmode=require"

	_, err := tasks.InitDB(connString)
	if err != nil {
		panic(err)
	}

	repo := tasks.NewSQLRepository(tasks.DB)

	args := os.Args[1:]

	// cli
	if len(args) > 0 {
		runCLI(repo, args)
		return
	}

	// server
	fmt.Println("Server HTTP attivo su http://localhost:8080")
	httpapi.Startserver(repo)
}
