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
		server.StartServer()

	} else if mode == "client" {
		client.StartClient()
	} else {
		fmt.Println("Wrong param")
	}

	// channel := make(chan string, 5)
	//
	// for x := 0; x < 5; x++ {
	// 	group.Add(1)
	// 	go asyncFunc(x, channel)
	// }
	//
	// group.Wait()
	// close(channel)
	//
	// for {
	// 	data, ok := <-channel
	// 	fmt.Println("OK", ok)
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Println(data)
	// }

}

// func asyncFunc(num int, c chan<- string) {
// 	defer group.Done()
// 	duration := time.Duration(rand.Intn(1000)) * time.Millisecond
// 	time.Sleep(duration)
//
// 	c <- fmt.Sprintln("EL NUMBER", num, "with duration:", duration)
// }
