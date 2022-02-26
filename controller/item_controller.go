package controller

import (
	"backend/service"
	"encoding/json"
	"net/http"
)

type IItemController interface {
	Handle(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Add(w http.ResponseWriter, r *http.Request)
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

func (c *ItemController) GetAll(w http.ResponseWriter, r *http.Request) {
	response, err := c.service.Items()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	json, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.Write(json)
}
func (c *ItemController) Add(w http.ResponseWriter, r *http.Request) {

}

func NewItemController(service service.IItemService) IItemController {
	return &ItemController{service: service}
}
