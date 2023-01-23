package router

import (
	"fmt"
	"net/http"
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/controllers"
	"search_engine_task/pkg/dbconn"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.IndentedJSON(http.StatusOK, "Alive")
	})

	router.GET("/ping", func(ctx *gin.Context) {
		dbconn.Dbconn()
		ctx.IndentedJSON(http.StatusOK, "Ping Successful")
	})	
	
	router.POST("/v1/savepage", controllers.SavePage)

	router.GET("/v1/getresult", controllers.GetResult)

	router.Run(":" + fmt.Sprint(config.Config.Server.Port))
}
