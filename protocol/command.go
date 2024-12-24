package protocol

import "net"

type Command struct {
	Name    string
	Handler func(conn net.Conn) error
}

type ClientCommand struct {
	Name    string
	Handler func(args string, conn net.Conn) (string, error)
}
