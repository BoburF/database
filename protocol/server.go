package protocol

import (
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	commands []Command
}

func (s *Server) create(host string, port int) (net.Listener, error) {
	return net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
}

func (s *Server) handleConnection(conn net.Conn) {

}

func (s *Server) Start(host string, port int) error {
	listener, err := s.create(host, port)
	if err != nil {
		return err
	}
	defer listener.Close()

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
	s.commands = append(s.commands, Command{
		Name:    strings.ToUpper(name),
		Handler: handler,
	})
}
