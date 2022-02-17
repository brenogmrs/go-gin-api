package main

import (
	"github.com/brenogmrs/go-gin-api/database"
	"github.com/brenogmrs/go-gin-api/server"
)

func main() {

	database.StartDB()

	server := server.NewServer()

	server.Run()
}
