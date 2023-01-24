package main

import (
	"context"
	"fmt"
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/controllers"
	"search_engine_task/pkg/external"
	"search_engine_task/pkg/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	config.LoadConfig()
	StartServer()

}

func StartServer() {
	router := gin.Default()
	mongoClient := external.NewMongoDB(Dbconn())
	searchservice := services.NewSearchService(mongoClient)
	ct := controllers.NewController(searchservice)
	ct.Routes(router.Group("/" + config.Config.Server.Version))
	router.Run(":" + fmt.Sprint(config.Config.Server.Port))
}

func Dbconn() (*mongo.Collection) {

	URI := config.Config.Database.Protocol + "://" + config.Config.Database.Host + ":" + fmt.Sprint(config.Config.Database.Port)

	clientOptions := options.Client().ApplyURI(URI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		panic(err)
	}

	return client.Database(config.Config.Database.DBName).Collection(config.Config.Database.Collection)
}
