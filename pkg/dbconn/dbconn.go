package dbconn

import (
	"context"
	"fmt"
	"os"
	"search_engine_task/cmd/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// type DBService interface {
// 	Dbconn()
// }

// type DB struct {
// 	col       *mongo.Collection
// 	dbservice DBService
// }

func Dbconn() *mongo.Collection {

	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)

	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {

		fmt.Println("Mongo.connect() ERROR: ", err)

		os.Exit(1)

	}

	// db.col = client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
}
