package protocol

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	Commands map[string]ClientCommand
}

func (c *Client) NewConnect(host string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	c.conn = conn

	return nil
}

func (c *Client) Call(command string, args string) (string, error) {
	handler, exists := c.Commands[command]
	if !exists {
		return "", errors.New("Command is not defined")
	}

	result, err := handler.Handler(args, c.conn)
	if err != nil {
		return "", err
	}

	return result, nil
}

func (c *Client) RegisterCommand(name string, handler func(args string, conn net.Conn) (string, error)) {
	c.Commands[name] = ClientCommand{
		Name:    strings.ToUpper(name),
		Handler: handler,
	}
}
