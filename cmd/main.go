package main

import (
	"search_engine_task/cmd/config"
	"search_engine_task/pkg/router"
)

func main() {
	config.LoadConfig()
	router.StartServer()
}