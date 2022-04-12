package routes

import (
	"awesomeProject/src/controllers"
	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("/category")
	{
		main.GET("/", controllers.ListCategories)
		main.GET("/:id", controllers.FindCategory)
		main.POST("/", controllers.CreateCategory)
		main.GET("/parms", controllers.ListParms)
	}

	return router
}
