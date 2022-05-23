package routes

import (
	"awesomeProject/src/controllers"
	"awesomeProject/src/repository"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, repo repository.Repository) *gin.Engine {
	handler := controllers.NewHandler(repo)
	main := router.Group("/category")
	{
		main.GET("/", handler.ListCategories)
		main.GET("/:id", handler.FindCategory)
		main.POST("/", handler.CreateCategory)
		main.GET("/parms", handler.ListParams)
	}
	parms := router.Group("/params")
	{
		parms.GET("/", handler.ListParams)
	}

	return router
}
