package main

import (
	"log"

	"github.com/BoburF/database/protocol"
)

func main() {
	server := protocol.Server{}
	if err := server.Start("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
}
