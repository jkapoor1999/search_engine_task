package main

import (
	"fmt"
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/controllers"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()
	StartServer()

}

func StartServer() {
	router := gin.Default()
	ct := controllers.Controller{}
	// ct.db.Dbconn()
	ct.Routes(router.Group("/v1"))
	router.Run(":" + fmt.Sprint(config.Config.Server.Port))
}
