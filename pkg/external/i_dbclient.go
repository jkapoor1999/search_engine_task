package external

import (
	"search_engine_task/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
)

type IDBClient interface {
	GetAllCollection() []bson.M
	InsertOnePage(context.Context, models.Page) (error)
}