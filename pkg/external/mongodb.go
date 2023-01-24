package external

import (
	"os"
	"search_engine_task/pkg/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type MongoDB struct {
	coll *mongo.Collection
}

func NewMongoDB(coll *mongo.Collection) *MongoDB {
	return &MongoDB{coll: coll}
}

func (c *MongoDB) GetAllCollection() []bson.M {
	cur, err := c.coll.Find(context.Background(), bson.D{})

	if err != nil {

		os.Exit(1)

	}

	var temp []bson.M

	for cur.Next(context.Background()) {

		var p bson.M

		err := cur.Decode(&p)

		if err != nil {

			os.Exit(1)

		}

		temp = append(temp, p)

	}

	defer cur.Close(context.Background())

	return temp
}

func (c *MongoDB) InsertOnePage(ctx context.Context, newPage models.Page) error {
	_, insertErr := c.coll.InsertOne(ctx, newPage)
	return insertErr
}
