package main

import (
	"backend/config"
	"backend/server"

	"log"
)

func main() {
	config := config.Get()

	err := server.NewServer().StartServer(config.ServerAddr)
	if err != nil {
		log.Fatalln(err)
	}
}
