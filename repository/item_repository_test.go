package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/model"
	"backend/repository"
)

func Test_GetAllItems(t *testing.T) {
	t.Run("should return all items correctly", func(t *testing.T) {
		itemsRepo, items := prepareForTest()
		want := items
		got, err := itemsRepo.GetAllItems()
		assert.Nil(t, err)
		assert.Equal(t, got, want)

	})

	t.Run("should return empty items if inMemoryDB is empty", func(t *testing.T) {
		items := []*model.Item{}

		itemsRepo := repository.NewItemRepository()

		want := items
		got, err := itemsRepo.GetAllItems()
		assert.Nil(t, err)
		assert.Equal(t, got, want)
	})
}

func Test_GettItem(t *testing.T) {
	t.Run("should get item successfully", func(t *testing.T) {
		itemsRepo, items := prepareForTest()
		want := items[0]
		got, err := itemsRepo.Get(items[0].Id)
		assert.Nil(t, err)
		assert.Equal(t, got, want)
	})
	t.Run("should return empty item if item doesnt exist with provided id", func(t *testing.T) {
		itemsRepo, _ := prepareForTest()
		want := &model.Item{}
		got, err := itemsRepo.Get(3)
		assert.Nil(t, err)
		assert.Equal(t, got, want)
	})
}

func Test_AddItem(t *testing.T) {
	t.Run("should add item successfully", func(t *testing.T) {
		itemsRepo, _ := prepareForTest()
		item := &model.Item{
			Id:   3,
			Text: "go to gym",
		}
		itemsRepo.NewItem(item)
		got, err := itemsRepo.Add(item)
		assert.Nil(t, err)
		assert.Equal(t, got, item)
	})
}

func prepareForTest() (repository.IItem, []*model.Item) {
	items := []*model.Item{
		{
			Id:   1,
			Text: "practice go",
		},
		{
			Id:   2,
			Text: "buy some cheese",
		},
	}
	// Create repository
	itemsRepo := repository.NewItemRepository()
	// Append items to repository
	for _, s := range items {
		itemsRepo.NewItem(s)
	}

	return itemsRepo, items
}
