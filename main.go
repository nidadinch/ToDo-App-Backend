package main

import (
	"backend/config"
	"backend/controller"
	"backend/repository"
	"backend/service"
	"fmt"
	"net/http"
)

func main() {
	config := config.Get()

	repository := repository.NewItemRepository()
	service := service.NewItemService(repository)
	controller := controller.NewItemController(service)

	http.HandleFunc("/", controller.Handle)
	err := http.ListenAndServe(config.ServerAddr, nil)
	if err != nil {
		fmt.Println(err)
	}

}
