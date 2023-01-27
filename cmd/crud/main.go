package main

import (
	config "CRUD/internal/app"
	logger "CRUD/internal/app/logs"
	"CRUD/internal/app/routes"
	"fmt"
)

func main() {
	fmt.Println(*config.GetConfig())
	logger.Init()
	routes.Init()
}
