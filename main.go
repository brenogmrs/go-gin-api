package main

import "github.com/brenogmrs/go-gin-api/server"

func main() {
	server := server.NewServer()

	server.Run()
}
