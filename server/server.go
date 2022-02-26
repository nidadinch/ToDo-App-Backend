package server

import (
	"backend/controller"
	"backend/repository"
	"backend/service"
	"fmt"
	"net/http"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) StartServer(port int) error {
	repository := repository.NewItemRepository()
	service := service.NewItemService(repository)
	controller := controller.NewItemController(service)

	http.HandleFunc("/", controller.Handle)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	return err
}
