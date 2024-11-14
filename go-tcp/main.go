package main

import (
	"fmt"
	"github.com/Maxi-Di-Mito/go-routines/client"
	"github.com/Maxi-Di-Mito/go-routines/server"
	"os"
)

// var group sync.WaitGroup

func main() {

	mode := os.Args[1]

	if mode == "server" {
		fmt.Println("Starting server")
		server.StartServer()
	} else if mode == "client" {
		fmt.Println("Starting client")
		client.StartClient()
	} else {
		fmt.Println("Wrong param")
	}
}
