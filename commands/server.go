package commands

import (
	"log"
	"net"
	"strings"

	"github.com/BoburF/database/protocol"
	"github.com/BoburF/database/storage"
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

	server.RegisterCommand("CREATE", func(conn net.Conn) error {
		result, err := CommandRead(conn)
		if err != nil {
			return err
		}

		data := strings.Split(result, " ")

		path := GeneratePath(data[0], GenerateTimestampID())

		err = storage.Create(path, data[1])
		if err != nil {
			return err
		}

		log.Println("Created ...")

		err = CommandResultWrite(conn, "SUCCESS")
		if err != nil {
			return err
		}

		return nil
	})
}
