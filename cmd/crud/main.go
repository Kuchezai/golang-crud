package main

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/routes"
)

func main() {
	config.GetConfig()
	logger.Init()
	routes.Init()
}
