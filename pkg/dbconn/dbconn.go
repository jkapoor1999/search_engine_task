package dbconn

import (
	"context"
	"fmt"
	"os"
	"search_engine_task/cmd/config"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Dbconn() *mongo.Collection {
	// const URI string = "mongodb://host.docker.internal:27017";
	// const URI string = "mongodb://mongo-container:27017"
	// const URI string = "mongodb://localhost:27017"

	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)

	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		fmt.Println("Mongo.connect() ERROR: ", err)

		os.Exit(1)

	}

	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
	// return client.Database("searchengine").Collection("pages")
}
