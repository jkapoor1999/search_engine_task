package dbconn

import (
	"context"
	"fmt"
	"search_engine_task/cmd/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbconn() (*mongo.Collection) {

	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)

	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
}
