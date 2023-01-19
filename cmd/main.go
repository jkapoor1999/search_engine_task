package main

import (
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/router"

	"github.com/gin-gonic/gin"
)

func main() {
	var r *gin.Engine
	config.LoadConfig()
	router.StartServer(r)

}
