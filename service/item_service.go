package service

import (
	"backend/model"
	"backend/repository"
)

type IItemService interface {
	Items() (*model.ItemsResponse, error)
}
type ItemService struct {
	Repository repository.IItem
}

func (s *ItemService) Items() (*model.ItemsResponse, error) {
	return nil, nil
}

func NewItemService(repository repository.IItem) IItemService {
	return &ItemService{Repository: repository}
}
