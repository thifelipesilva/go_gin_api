package main

import (
	"github.com/thifelipesilva/go_gin_api/database"
	"github.com/thifelipesilva/go_gin_api/routes"
)

func main() {
	database.ConecctionWithDB()
	routes.HandleRequests()
}
