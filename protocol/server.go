package protocol

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

type Server struct {
	commands map[string]Command
}

func (s *Server) create(host string, port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
}

func (s *Server) handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		lengthOfCommand := make([]byte, 2, 2)
		_, err := conn.Read(lengthOfCommand)
		if err != nil {
			log.Panic(err)
			break
		}

		lengthOfCommandParsed, err := strconv.Atoi(string(lengthOfCommand))
		if err != nil || lengthOfCommandParsed <= 0 {
			log.Panic(err)
			break
		}

		command := make([]byte, lengthOfCommandParsed)
		_, err = conn.Read(command)
		if err != nil {
			log.Panic(err)
			break
		}

		checkByte := make([]byte, 1)
		_, err = conn.Read(checkByte)
		if err != nil {
			log.Panic(err)
			break
		}

		if !bytes.Equal(checkByte, []byte{'\x00'}) {
			log.Println("Invalid check byte received:", checkByte)
			return
		}

		commandParsed := strings.ToUpper(string(command))

		handler, exists := s.commands[commandParsed]
		if !exists {
			log.Println("Unknown command:", commandParsed)
			return
		}
		handler.Handler(conn)
	}
}

func (s *Server) Start(host string, port int) error {
	listener, err := s.create(host, port)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.commands = make(map[string]Command)

	log.Println("Server started at port:", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		go s.handleConnection(conn)
	}
}

func (s *Server) RegisterCommand(name string, handler func(conn net.Conn) error) {
	s.commands[name] = Command{
		Name:    strings.ToUpper(name),
		Handler: handler,
	}
}
