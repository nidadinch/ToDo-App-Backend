package main

import (
	"backend/server"

	"log"
)

func main() {

	err := server.NewServer().StartServer(3000)
	if err != nil {
		log.Fatalln(err)
	}
}
