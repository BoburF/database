package commands

import (
	"net"

	"github.com/BoburF/database/protocol"
)

func RegisterPredefinedClientCommands(client *protocol.Client) {
	client.RegisterCommand("PING", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "PING", args)
		if err != nil {
			return "", err
		}
		result, err := CommandResultRead(conn)
		return result, err
	})

	client.RegisterCommand("ECHO", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "ECHO", args)
		result, err := CommandResultRead(conn)
		return result, err

	})

	client.RegisterCommand("QUIT", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "QUIT", args)
		return "", err
	})

	client.RegisterCommand("CREATE", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "CREATE", args)
		result, err := CommandResultRead(conn)
		return result, err
	})

    client.RegisterCommand("GET", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "GET", args)
		result, err := CommandResultRead(conn)
		return result, err
	})
}
