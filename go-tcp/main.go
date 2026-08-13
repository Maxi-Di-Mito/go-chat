package main

import (
	"fmt"
	"os"

	"github.com/Maxi-Di-Mito/go-routines/client"
	"github.com/Maxi-Di-Mito/go-routines/server"
)

// var group sync.WaitGroup

func main() {

	userName := ""
	mode := os.Args[1]
	if len(os.Args) == 3 {
		userName = os.Args[2]
	}

	switch mode {
	case "server":
		fmt.Println("Starting server")
		server.StartServer()
	case "client":
		fmt.Println("Starting client")
		client.StartClient(userName)
	default:
		fmt.Println("Wrong param")
	}
}
