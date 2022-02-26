package controller_test

import (
	"backend/controller"
	"backend/mock"
	"backend/model"
	"encoding/json"
	"errors"
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

		resBody, _ := json.Marshal(serviceReturn)

		assert.Equal(t, string(resBody), w.Body.String())
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})
	t.Run("should return error if service fails", func(t *testing.T) {
		service := mock.NewMockIItemService(gomock.NewController(t))
		serviceErr := errors.New("test err")

		service.EXPECT().Items().Return(nil, serviceErr)
		controller := controller.NewItemController(service)

		r := httptest.NewRequest(http.MethodGet, "/items", nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}

func Test_Add(t *testing.T) {
	t.Run("should add new item successfully", func(t *testing.T) {
		service := mock.NewMockIItemService(gomock.NewController(t))
		serviceReturn := &model.ItemsResponse{1: "buy some milk"}

		service.EXPECT().Add("buy some milk").Return(serviceReturn, nil)
		controller := controller.NewItemController(service)

		r := httptest.NewRequest(http.MethodPost, "/item", nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		resBody, _ := json.Marshal(serviceReturn)

		assert.Equal(t, string(resBody), w.Body.String())
		assert.Equal(t, w.Result().StatusCode, http.StatusOK)
		assert.Equal(t, "application/json", w.Header().Get("content-type"))
	})
	t.Run("should return error if service fails", func(t *testing.T) {
		service := mock.NewMockIItemService(gomock.NewController(t))
		serviceErr := errors.New("test err")

		service.EXPECT().Add("buy some milk").Return(nil, serviceErr)
		controller := controller.NewItemController(service)

		r := httptest.NewRequest(http.MethodGet, "/items", nil)
		w := httptest.NewRecorder()
		controller.Handle(w, r)

		assert.Equal(t, w.Result().StatusCode, http.StatusInternalServerError)
		assert.Equal(t, w.Body.String(), serviceErr.Error())
	})
}
