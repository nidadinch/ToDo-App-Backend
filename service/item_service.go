package service

import (
	"backend/model"
	"backend/repository"
)

type IItemService interface {
	Items() (*model.ItemsResponse, error)
	Add(text string) (*model.ItemsResponse, error)
}
type ItemService struct {
	Repository repository.IItem
}

func (s *ItemService) Items() (*model.ItemsResponse, error) {
	items, err := s.Repository.GetAllItems()
	m := model.ItemsResponse{}

	for _, v := range items {
		m[v.Id] = v.Text
	}

	return &m, err
}

func (s *ItemService) Add(text string) (*model.ItemsResponse, error) {
	return nil, nil
}

func NewItemService(repository repository.IItem) IItemService {
	return &ItemService{Repository: repository}
}
