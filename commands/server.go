package commands

import (
	"net"

	"github.com/BoburF/database/protocol"
)

func RegisterPredefinedCommands(server *protocol.Server) {
	server.RegisterCommand("PING", func(conn net.Conn) error {
		_, err := CommandRead(conn)
		if err != nil {
			return err
		}
		err = CommandResultWrite(conn, "PONG")
		return err
	})

	server.RegisterCommand("ECHO", func(conn net.Conn) error {
		result, err := CommandRead(conn)
		if err != nil {
			return err
		}
		err = CommandResultWrite(conn, result)
		return err
	})

	server.RegisterCommand("QUIT", func(conn net.Conn) error {
		err := CommandResultWrite(conn, "BYE:)")
		conn.Close()
		return err
	})
}
