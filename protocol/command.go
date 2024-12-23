package protocol

import "net"

type Command struct {
	Name    string
	Handler func(conn net.Conn) error
}

func RegisterPredefinedCommands(server *Server) {
	server.RegisterCommand("PING", func(conn net.Conn) error {
		_, err := conn.Write([]byte("PONG\n"))
		return err
	})

	server.RegisterCommand("ECHO", func(conn net.Conn) error {
		_, err := conn.Write([]byte("ECHO received\n"))
		return err
	})

	server.RegisterCommand("QUIT", func(conn net.Conn) error {
		_, err := conn.Write([]byte("Goodbye\n"))
		conn.Close()
		return err
	})
}
