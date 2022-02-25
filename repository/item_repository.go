package repository

import (
	"backend/model"
)

type Item struct {
	items []*model.Item
}

type IItem interface {
	NewItem(item *model.Item)
	GetAllItems() ([]*model.Item, error)
	Get(id int) (*model.Item, error)
}

func (i *Item) NewItem(item *model.Item) {
	i.items = append(i.items, item)
}

func (i *Item) GetAllItems() ([]*model.Item, error) {
	return i.items, nil
}

func (i *Item) Get(id int) (*model.Item, error) {
	for _, item := range i.items {
		if item.Id == id {
			return item, nil
		}
	}
	empty := &model.Item{}
	return empty, nil
}

func NewItemRepository() IItem {
	DB := []*model.Item{}
	return &Item{items: DB}
}
