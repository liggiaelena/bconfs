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
	c.JSON(
		200,
		gin.H{
			"categorias:": category,
		},
	)
}

func FindCategory(c *gin.Context) {
	id := c.Param("id")

	db := database.GetDatabase()
	var category models.Category
	err := db.Find(&category, id).Error

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, category)
}

func CreateCategory(c *gin.Context) {
	db := database.GetDatabase()

	var category models.Category
	
}
