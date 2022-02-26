package controller_test

import (
	"backend/controller"
	"backend/mock"
	"backend/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAll(t *testing.T) {
	t.Run("should return items correctly", func(t *testing.T) {
		service := mock.NewMockIItemService(gomock.NewController(t))
		serviceReturn := &model.ItemsResponse{1: "buy some milk", 2: "go to gym"}

		service.EXPECT().Items().Return(serviceReturn, nil)
		controller := controller.NewItemController(service)

		r := httptest.NewRequest(http.MethodGet, "/items", nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		actual := serviceReturn
		json.Unmarshal(w.Body.Bytes(), actual)

		assert.Equal(t, serviceReturn, actual)
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})
}
