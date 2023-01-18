package main

import (
	config "CRUD/internal/app"
	"CRUD/internal/app/routes"
)

func main() {
	config.Load()
	routes.Init()
}
