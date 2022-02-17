package server

import (
	"log"

	"github.com/brenogmrs/go-gin-api/server/routes"
	"github.com/gin-gonic/gin"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "5000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("Server is running on port", s.port)
	log.Fatal(router.Run(":" + s.port))
}
