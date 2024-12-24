package protocol

import "net"

type Command struct {
	Name    string
	Handler func(conn net.Conn) error
}
