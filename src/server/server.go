package server

import (
	"awesomeProject/src/repository"
	"awesomeProject/src/routes"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	Port   string
	Server *gin.Engine
}

func NewServer() Server {
	return Server{
		Port:   "4000",
		Server: gin.Default(),
	}
}

func (s *Server) Run() {
	repo := repository.DefaultRepository()
	router := routes.ConfigRoutes(s.Server, repo)

	log.Println("server is running at port" + s.Port)
	log.Fatal(router.Run(":" + s.Port))
	//impede que a aplicacao morra, mas faz ela ficar esperando request
}
