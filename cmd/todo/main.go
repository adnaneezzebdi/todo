package main

import (
	"os"
	"todo/internal/cli"
	"todo/internal/httpapi"
)

func main() {

	if len(os.Args) < 2 {
		println("usage: todo [cli|server]")
		return
	}

	switch os.Args[1] {
	case "cli":
		cli.Run()

	case "server":
		httpapi.Startserver()

	default:
		println("comando sconosciuto:", os.Args[1])
	}
}
