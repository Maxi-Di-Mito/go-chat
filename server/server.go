package server

import (
	"fmt"
	"net"
)

func StartServer() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	connectionListener(listener)

}

func connectionListener(listener net.Listener) {
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("ERROR LISTENER", err)
			continue
		}
		go receiver(conn)
	}
}

func receiver(conn net.Conn) {
	defer conn.Close()

	in := []byte{}

	_, err := conn.Read(in)
	if err != nil {
		fmt.Println("ERRO", err)
		return
	}

	fmt.Println("MESSAGE", in)
}
