package commands

import (
	"fmt"
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

func RegisterPredefinedClientCommands(client *protocol.Client) {
	client.RegisterCommand("PING", func(args string, conn net.Conn) error {
		_, err := conn.Write([]byte(fmt.Sprintf("%02dPING\x00%s", 4, args)))
		return err
	})

	client.RegisterCommand("ECHO", func(args string, conn net.Conn) error {
		_, err := conn.Write([]byte(fmt.Sprintf("%02dECHO\x00%s", 4, args)))
		return err
	})

	client.RegisterCommand("QUIT", func(args string, conn net.Conn) error {
		_, err := conn.Write([]byte(fmt.Sprintf("%02dQUIT\x00%s", 4, args)))
		conn.Close()
		return err
	})
}
