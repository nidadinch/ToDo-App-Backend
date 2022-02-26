package controller

import (
	"backend/service"
	"net/http"
)

type IItemController interface {
	Handle(w http.ResponseWriter, r *http.Request)
}

type ItemController struct {
	service service.IItemService
}

func (c *ItemController) Handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/items" {
			c.GetAll(w, r)
		}
	case http.MethodPost:
		if r.URL.Path == "/item" {
			c.Add(w, r)
		}
	}

}

func NewItemController(service service.IItemService) IItemController {
	return &ItemController{service: service}
}
