package service_test

import (
	"backend/mock"
	"backend/model"
	"backend/service"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func CreateMockRepository(t *testing.T) *mock.MockIItem {
	return mock.NewMockIItem(gomock.NewController(t))
}

func Test_GetAll(t *testing.T) {
	mockRepo := CreateMockRepository(t)
	mockItems := []*model.Item{
		{Id: 1, Text: "practice go"},
		{Id: 2, Text: "go to gym"},
	}

	mockRepo.EXPECT().GetAllItems().Return(mockItems, nil)
	itemService := service.NewItemService(mockRepo)
	items, err := itemService.Items()

	assert.Equal(t, &model.ItemsResponse{1: "practice go", 2: "go to gym"}, items)
	assert.Nil(t, err)
}

func Test_Add(t *testing.T) {
	t.Run("should add todo item with id=1 when there is any items inside", func(t *testing.T) {
		mockRepo := CreateMockRepository(t)
		text := "buy some milk"

		mockItem := &model.Item{Id: 1, Text: text}

		mockRepo.EXPECT().GetAllItems().Return([]*model.Item{}, nil)
		mockRepo.EXPECT().Add(mockItem).Return(mockItem, nil)

		itemService := service.NewItemService(mockRepo)
		item, err := itemService.Add(text)

		assert.Equal(t, &model.ItemsResponse{1: text}, item)
		assert.Nil(t, err)
	})
	t.Run("should add todo item with id = 2 when there is 1 item inside", func(t *testing.T) {
		mockRepo := CreateMockRepository(t)
		text := "buy some milk"
		text2 := "go to gym"

		mockItem := []*model.Item{{Id: 1, Text: text}}
		mockItem2 := &model.Item{Id: 2, Text: text2}
		mockRepo.EXPECT().GetAllItems().Return(mockItem, nil)
		mockRepo.EXPECT().Add(mockItem2).Return(mockItem2, nil)

		itemService := service.NewItemService(mockRepo)
		item, err := itemService.Add(text2)

		assert.Equal(t, &model.ItemsResponse{2: text2}, item)
		assert.Nil(t, err)
	})

}
