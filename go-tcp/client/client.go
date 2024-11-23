package client

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
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
		inputChannel <- strings.TrimSpace(input)
	}
}

func SendInput(in string) {
	inputChannel <- in
}

func processInput(conn net.Conn) {
	for {
		input := <-inputChannel
		if inputIsAction(input) {
			makeAction(input)
		} else {
			conn.Write([]byte(input))
		}
	}
}

func inputIsAction(input string) bool {
	if input == "EXIT" {
		return true
	}
	return false
}

func makeAction(action string) {
	toPrint <- fmt.Sprintf("Executing action: %s", action)
	if action == "EXIT" {
		close(toPrint)
		close(inputChannel)
		group.Done()
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

		toPrint <- string(input[:len])
	}
}

func printer() {
	for {
		data, ok := <-toPrint
		if !ok {
			break
		}
		fmt.Println("===============================================================\n", data)
	}
	group.Done()
}
