package controllers

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"github.com/gin-gonic/gin"
)

func ListCategories(c *gin.Context) {

	db := database.GetDatabase()
	var category []models.Category
	err := db.Find(&category).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	//c.JSON(
	//	200,
	//	gin.H{
	//		category,
	//	},
	//)
	c.JSON(200, category)
}

//db := database.GetDatabase()
//var p []models.Book
//err := db.Find(&p).Error
//
//if err != nil {
//c.JSON(400, gin.H{
//"error": "cannot find product by id: " + err.Error(),
//})
//return
//}
//
//c.JSON(200, p)
