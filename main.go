package main

import (
	"fmt"
	"log"
	"time"

	"github.com/BoburF/database/commands"
	"github.com/BoburF/database/protocol"
	"github.com/BoburF/database/storageformat"
)

type Go struct {
	Name string `bsf:"go_name"`
}

type MessageEcho struct {
	Name string `json:"nam" bsf:"nam"`
	Mess string `json:"message" bsf:"message"`
	Go
}

type Data struct {
	Name    string `bsf:"name"`
	SurName string `bsf:"surname"`
	Phone   string `bsf:"phone"`
}

func main() {
	go startServer()

	time.Sleep(time.Second * 2)
	client := protocol.Client{Commands: make(map[string]protocol.ClientCommand)}
	if err := client.NewConnect("localhost", 8080); err != nil {
		log.Println("Error starting server:", err)
	}
	commands.RegisterPredefinedClientCommands(&client)

	data := Data{
		Name:    "Bobur",
		SurName: "Abdullayev",
		Phone:   "998939752577",
	}
	collection := "maker"

	query := fmt.Sprintf("%s %s", collection, string(storageformat.ToStorageFormat(data)))

	result, err := client.Call("CREATE", query)
	if err != nil {
		log.Println("Error calling command", err)
	}

	log.Println("Result:", result)

	result, err = client.Call("GET", collection+" id "+result)
	if err != nil {
		log.Println("Error calling command", err)
	}

	log.Println("Result:", result)

	result, err = client.Call("GETALL", collection)
	if err != nil {
		log.Println("Error calling command", err)
	}

	log.Println("Result:", result)

	result, err = client.Call("QUIT", "")
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
