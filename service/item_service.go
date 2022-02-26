package service

import (
	"backend/model"
	"backend/repository"
)

type IItemService interface {
	Items() (*model.ItemsResponse, error)
	Add(text string) (*model.Item, error)
}
type ItemService struct {
	Repository repository.IItem
}

func (s *ItemService) Items() (*model.ItemsResponse, error) {
	items, err := s.Repository.GetAllItems()
	m := model.ItemsResponse{}

	for _, v := range items {
		m = append(m, *v)
	}
	return &m, err
}

func (s *ItemService) Add(text string) (*model.Item, error) {
	items, _ := s.Repository.GetAllItems()
	itemId := len(items) + 1
	itemToAdd := &model.Item{Id: itemId, Text: text}
	item, err := s.Repository.Add(itemToAdd)

	m := model.Item{}
	m = *item

	return &m, err
}

func NewItemService(repository repository.IItem) IItemService {
	return &ItemService{Repository: repository}
}
