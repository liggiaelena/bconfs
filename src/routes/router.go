package routes

import (
	"awesomeProject/src/controllers"
	"awesomeProject/src/repository"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine, repo repository.Repository) *gin.Engine {
	handler := controllers.NewHandler(repo)

	router.GET("/category", handler.ListCategories)
	router.GET("/category/:id", handler.FindCategory)
	router.POST("/category", handler.CreateCategory)

	router.GET("/params", handler.ListParams)

	return router
}
