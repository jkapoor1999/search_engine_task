package services

import (
	"context"
	"search_engine_task/mocks"
	"search_engine_task/pkg/models"
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/mock"
)

func TestNewSearchService(t *testing.T) {

}

func TestSearchService_SavePage(t *testing.T) {

	mockService := mocks.NewIDBClient(t)
	mockService.On("InsertOnePage", mock.Anything, mock.Anything).Return(nil)
	searchservice := NewSearchService(mockService)

	input := models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	var ctx context.Context
	err := searchservice.SavePage(ctx, input)

	assert.Equal(t, nil, err)
}

func TestSearchService_GetResult(t *testing.T) {
	mockService := mocks.NewIDBClient(t)
	mockService.On("GetAllCollection", mock.Anything, mock.Anything).Return(nil)
	searchservice := NewSearchService(mockService)

	input := models.Page{
		Title:    "Page 10",
		Keywords: []string{"wrd1", "wrd2"},
	}
	var ctx context.Context
	err := searchservice.SavePage(ctx, input)

	assert.Equal(t, nil, err)
}
