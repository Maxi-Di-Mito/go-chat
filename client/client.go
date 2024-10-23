package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"sync"
)

var group sync.WaitGroup

var inputChannel chan string
var toPrint chan string

func StartClient() {
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	toPrint = make(chan string, 20)
	inputChannel = make(chan string)

	go receiver(conn)
	group.Add(1)
	go printer()

	go inputer()
	go processInput(conn)

	group.Wait()
}

func inputer() {
	reader := bufio.NewReader(os.Stdin)
	for {
		input, _ := reader.ReadString('\n')
		inputChannel <- input
	}
}

func processInput(conn net.Conn) {
	for {
		input := <-inputChannel

		conn.Write([]byte(input))
	}
}

func receiver(conn net.Conn) {
	input := make([]byte, 1024)
	for {
		len, err := conn.Read(input)
		if err != nil {
			panic(err)
		}
		if len == 0 {
			continue
		}

		toPrint <- string(input)
	}
}

func printer() {
	for {
		data, ok := <-toPrint
		if !ok {
			break
		}
		fmt.Println("==============================\n", data)
	}
	group.Done()
}
