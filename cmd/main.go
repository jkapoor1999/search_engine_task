package main

import (
	"fmt"
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/controllers"
	"search_engine_task/pkg/dbconn"
	"search_engine_task/pkg/external"
	"search_engine_task/pkg/services"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	StartServer()

}

func StartServer() {
	router := gin.Default()
	mongoClient := external.NewMongoDB(dbconn.Dbconn())
	searchservice := services.NewSearchService(mongoClient)
	ct := controllers.NewController(searchservice)
	ct.Routes(router.Group("/" + config.Config.Server.Version))
	router.Run(":" + fmt.Sprint(config.Config.Server.Port))
}
