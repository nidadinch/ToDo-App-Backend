package repository_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"backend/model"
	"backend/repository"
)

func Test_GetAllItems(t *testing.T) {
	t.Run("should return all items correctly", func(t *testing.T) {
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
