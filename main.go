package main

import (
	"log"

	"github.com/BoburF/database/commands"
	"github.com/BoburF/database/protocol"
)

func main() {
	server := protocol.Server{}
	if err := server.Start("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
	commands.RegisterPredefinedCommands(&server)

	client := protocol.Client{}
	if err := client.NewConnect("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
	commands.RegisterPredefinedClientCommands(&client)

	client.Call("PING", "")
}
