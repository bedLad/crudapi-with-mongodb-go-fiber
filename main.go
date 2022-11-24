package main

import (
	"fmt"

	"github.com/bedLad/go-fiber-mongo-hrms/database"
	"github.com/bedLad/go-fiber-mongo-hrms/routes"
)

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017"

func main() {
	// connect to the database
	database.Connect(mongoURI, dbName)

	// start the server
	fmt.Println("Starting server @ PORT:3000")
	// routes and server start at 3000
	routes.Init()
}
