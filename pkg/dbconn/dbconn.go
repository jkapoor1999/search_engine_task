package dbconn

import (
	"context"
	"fmt"
	"os"
	"search_engine_task/cmd/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoService struct {
	collection *mongo.Collection
}

func NewMongoService(collection *mongo.Collection) DB {
	return &MongoService{collection: collection}
}



func Dbconn() *mongo.Collection {

	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)

	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		fmt.Println("Mongo.connect() ERROR: ", err)

		os.Exit(1)

	}

	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)

}
