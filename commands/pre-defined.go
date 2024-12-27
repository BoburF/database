package commands

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/BoburF/database/protocol"
)

func RegisterPredefinedCommands(server *protocol.Server) {
	server.RegisterCommand("PING", func(conn net.Conn) error {
		err := CommandResultWrite(conn, "PONG")
		return err
	})

	server.RegisterCommand("ECHO", func(conn net.Conn) error {
		result, err := CommandResultRead(conn)
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
		return "", err
	})

	client.RegisterCommand("QUIT", func(args string, conn net.Conn) (string, error) {
		err := CommandWrite(conn, "QUIT", args)
		return "", err
	})
}

func CommandWrite(conn net.Conn, command string, args string) error {
	_, err := conn.Write([]byte(fmt.Sprintf("%02d%s\x00%s", len(command), command, args)))
	return err
}

func CommandResultWrite(conn net.Conn, result string) error {
	_, err := conn.Write([]byte(fmt.Sprintf("%d\x00%s", len(result), result)))
	return err
}

func CommandResultRead(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)

	lengthStr, err := reader.ReadString('\x00')
	if err != nil {
		return "", fmt.Errorf("failed to read length prefix: %w", err)
	}

	lengthStr = strings.TrimSuffix(lengthStr, "\x00")
	length, err := strconv.Atoi(lengthStr)
	if err != nil {
		return "", fmt.Errorf("failed to parse length: %w", err)
	}

	resultBuffer := make([]byte, length)
	_, err = reader.Read(resultBuffer)
	if err != nil {
		return "", fmt.Errorf("failed to read result data: %w", err)
	}

	return string(resultBuffer), nil
}
