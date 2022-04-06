package server

import (
	"awesomeProject/src/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port   string
	server *gin.Engine
}

func NewServer() Server {
	return Server{
		port:   "4000",
		server: gin.Default(),
	}
}

func (s *Server) Run() {
	router := routes.ConfigRoutes(s.server)

	log.Println("server is running at port" + s.port)
	log.Fatal(router.Run(":" + s.port))
}
