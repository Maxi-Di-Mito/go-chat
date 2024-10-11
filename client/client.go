package client

import "net"

func StartClient() {
	conn, err := net.Dial("tcp", "::8080")
	if err != nil {
		panic(err)
	}
}
