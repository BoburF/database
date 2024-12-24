package commands

import (
	"net"

	"github.com/BoburF/database/protocol"
)

func RegisterPredefinedCommands(server *protocol.Server) {
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
