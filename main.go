package main

import (
	"log"
	"time"

	"github.com/BoburF/database/commands"
	"github.com/BoburF/database/protocol"
)

func main() {
	go startServer()

	time.Sleep(time.Second * 2)
	client := protocol.Client{Commands: make(map[string]protocol.ClientCommand)}
	if err := client.NewConnect("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
	commands.RegisterPredefinedClientCommands(&client)

	result, err := client.Call("PING", "")
	if err != nil {
		log.Println("Error calling command")
	}

	log.Println("Result:", result)

	result, err = client.Call("ECHO", "Bobur zo'r bola")
	if err != nil {
		log.Println("Error calling command")
	}

	log.Println("Result:", result)
}

func startServer() {
	server := protocol.Server{
		Commands: make(map[string]protocol.Command),
	}
	commands.RegisterPredefinedCommands(&server)
	if err := server.Start("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
	log.Println("=========", server.Commands)
}
