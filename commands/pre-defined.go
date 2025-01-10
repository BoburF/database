package commands

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func CommandWrite(conn net.Conn, command string, args string) error {
	_, err := conn.Write([]byte(fmt.Sprintf("%02d%s\x00%d\x00%s", len(command), command, len(args), args)))
	return err
}

func CommandRead(conn net.Conn) (string, error) {
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
