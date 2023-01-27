package external

import (
	"search_engine_task/pkg/models"

	"golang.org/x/net/context"
)

type IDBClient interface {
	GetAllCollection() []models.Page
	InsertOnePage(context.Context, models.Page) error
}
