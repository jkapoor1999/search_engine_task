package services

import (
	"context"
	"search_engine_task/pkg/models"
)

type ISearchService interface {
	SavePage(context.Context, models.Page) error
	GetResult(context.Context, models.Keywords) ([]string, error)
}
