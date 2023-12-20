package main

import (
	"github.com/rafawildfire42/go-gin-api-study/database"
	"github.com/rafawildfire42/go-gin-api-study/routes"
)

func main() {
	database.ConnectDatabase()
	routes.HandleRequets()
}
