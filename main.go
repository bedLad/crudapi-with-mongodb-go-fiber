package main

import (
	"fmt"

	"github.com/bedLad/go-fiber-mongo-hrms/database"
	"github.com/bedLad/go-fiber-mongo-hrms/routes"
)

const dbName = "fiber-hrms"
const mongoURI = "mongodb+srv://rutvik:rUTVIK%4013@golang-coll.zmxavsk.mongodb.net/?retryWrites=true&w=majority"

func main() {
	// connect to the database
	database.Connect(mongoURI, dbName)

	// start the server
	fmt.Println("Starting server @ PORT:3000")
	// routes and server start at 3000
	routes.Init()
}
