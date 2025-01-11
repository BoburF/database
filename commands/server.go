package commands

import (
	"net"
	"os"
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

		fileId := GenerateTimestampID() + ".bsf"
		path := GeneratePath(data[0], fileId)

		err = storage.Create(path, data[1])
		if err != nil {
			return err
		}

		pathMeta := GeneratePath(data[0], "meta")
		file, err := os.OpenFile(pathMeta, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = file.WriteString(fileId + "\n")
		if err != nil {
			return err
		}

		err = CommandResultWrite(conn, fileId[:len(fileId)-4])
		if err != nil {
			return err
		}

		return nil
	})

	server.RegisterCommand("GET", func(conn net.Conn) error {
		result, err := CommandRead(conn)
		if err != nil {
			return err
		}

		data := strings.Split(result, " ")

		if data[1] == "id" {
			path := GeneratePath(data[0], data[2]+".bsf")

			result, err = storage.Read(path)
			if err != nil {
				return err
			}

			err = CommandResultWrite(conn, result)
			if err != nil {
				return err
			}
		} else {
			err = CommandResultWrite(conn, "Other indexes are not supported")
			if err != nil {
				return err
			}
		}

		return nil
	})

	server.RegisterCommand("GETALL", func(conn net.Conn) error {
		result, err := CommandRead(conn)
		if err != nil {
			return err
		}

		data := strings.Split(result, " ")

		path := GeneratePath(data[0], "meta")

		result, err = storage.Read(path)
		if err != nil {
			return err
		}

		err = CommandResultWrite(conn, result)
		if err != nil {
			return err
		}

		return nil
	})
}
