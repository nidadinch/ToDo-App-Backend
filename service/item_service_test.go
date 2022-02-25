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
