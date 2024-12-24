package protocol

import (
	"errors"
	"fmt"
	"net"
	"strings"
)

type Client struct {
	conn     net.Conn
	commands map[string]ClientCommand
}

func (c *Client) NewConnect(host string, port int) error {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return err
	}
	c.conn = conn
	c.commands = make(map[string]ClientCommand)

	return nil
}

func (c *Client) Call(command string, args string) error {
	handler, exists := c.commands[command]
	if !exists {
		return errors.New("Command is not defined")
	}

	handler.Handler(args, c.conn)

	return nil
}

func (c *Client) RegisterCommand(name string, handler func(args string, conn net.Conn) error) {
	c.commands[name] = ClientCommand{
		Name:    strings.ToUpper(name),
		Handler: handler,
	}
}
