package controllers

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"search_engine_task/mocks"
	"search_engine_task/pkg/models"
	"search_engine_task/pkg/services"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_Get(t *testing.T) {
	x := mocks.NewIDBClient(t)
	x.On("GetAllCollection").Return(nil)
	ss := services.NewSearchService(x)
	controller := NewController(ss)
	router := gin.Default()
	router.GET("/get", controller.Get)
	input := &models.Keywords{
		User_keywords: []string{"wrd1", "wrd2"},
	}
	jsonInput, _ := json.Marshal(input)
	req := httptest.NewRequest("GET", "/get", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}

func TestController_Insert(t *testing.T) {
	mockService := mocks.NewISearchService(t)
	mockService.On("SavePage", mock.Anything, mock.Anything).Return(nil)
	controller := NewController(mockService)

	router := gin.Default()
	router.POST("/insert", controller.Insert)
	input := &models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	jsonInput, _ := json.Marshal(input)
	req := httptest.NewRequest("POST", "/insert", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusCreated, resp.Code)

}

func TestController_Check(t *testing.T) {
	mockService := mocks.NewISearchService(t)
	controller := NewController(mockService)

	router := gin.Default()
	router.GET("/", controller.Check)
	input := &models.Page{}		
	jsonInput, _ := json.Marshal(input)
	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(jsonInput))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	assert.Equal(t, http.StatusOK, resp.Code)
}
