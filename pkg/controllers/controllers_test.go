package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"search_engine_task/mocks"
	"search_engine_task/pkg/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestController_Get(t *testing.T) {

	mockService := mocks.NewISearchService(t)
	mockService.On("GetResult", mock.Anything, mock.Anything).Return([]string{}, nil)
	controller := NewController(mockService)
	gin.SetMode(gin.TestMode)
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
	assert.NotEmpty(t, resp.Body)

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
	assert.NotEmpty(t, resp.Body)

}

func TestController_Check(t *testing.T) {
	mockService := mocks.NewISearchService(t)
	controller := NewController(mockService)
	router := gin.Default()
	router.GET("/", controller.Check)
	req := httptest.NewRequest("GET", "/", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)	
	responseData, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, "\"Alive\"", string(responseData))
	assert.Equal(t, http.StatusOK, resp.Code)
}

// func TestController_UseComputeResultError1(t *testing.T) {
//     gin.SetMode(gin.TestMode)
//     router := gin.Default()
//     mockSearchService := mocks.NewISearchService(t)
//     mockController := NewController(mockSearchService)
//     test := `"dkf ldjf ldlfjldj ljf"`
//     jsonBytes, _ := json.Marshal(test)
//     router.GET("/compute", mockController.UseComputeResult)
//     req := httptest.NewRequest("GET", "/compute", bytes.NewBuffer([]byte(jsonBytes)))
//     resp := httptest.NewRecorder()
//     router.ServeHTTP(resp, req)
//     assert.Equal(t, http.StatusBadRequest, resp.Code)
//     assert.NotEmpty(t, resp.Body)
// }
// func TestController_UseComputeResultError2(t *testing.T) {
//     gin.SetMode(gin.TestMode)
//     router := gin.Default()
//     mockSearchService := mocks.NewISearchService(t)
//     mockController := NewController(mockSearchService)
//     testPage := models.Keywords{ArrayOfString: []string{"x", "y"}}
//     jsonBytes, _ := json.Marshal(testPage)
//     mockSearchService.On("ComputeResult", mock.Anything, mock.Anything).Return([]string{}, errors.New("error"))
//     router.GET("/compute", mockController.GetRes)
//     req := httptest.NewRequest("GET", "/compute", bytes.NewBuffer([]byte(jsonBytes)))
//     resp := httptest.NewRecorder()
//     router.ServeHTTP(resp, req)
//     assert.Equal(t, http.StatusInternalServerError, resp.Code)
//     assert.NotEmpty(t, resp.Body)
// }