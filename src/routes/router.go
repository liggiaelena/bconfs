package routes

import (
	"awesomeProject/src/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/category")
	{
		main.GET("/", controllers.ListCategories)
	}
	return router
}
