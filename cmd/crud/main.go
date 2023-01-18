package main

import (
	config "CRUD/internal/app"
	"CRUD/internal/app/routes"
	"fmt"
)

func main() {
	configuration := config.Load()
	fmt.Println(configuration)

	routes.Init()
}
