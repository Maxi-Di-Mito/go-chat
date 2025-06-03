package server

import (
	"net"
)

// ///////////////////////////////////////////////////////////////
// Client Object and logic
type Client struct {
	id   string
	name string
	conn net.Conn
}
