package controllers

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"fmt"
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

	if category.ID == 0 {
		c.JSON(404, gin.H{
			"error": "this category does not exist",
		})
		return
	}

	c.JSON(200, category)
}

func CreateCategory(c *gin.Context) {
	db := database.GetDatabase()

	var category models.Category

	err := c.ShouldBindJSON(&category) //how get the body?
	fmt.Println(category)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	db.Find(&category, "name = ?", category.Name)
	fmt.Println(category)
	if category.ID != 0 {
		c.JSON(409, gin.H{
			"error": "This category already exist",
		})
		return
	}

	err = db.Create(&category).Error
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot create category: " + err.Error(),
		})
		return
	}

	c.JSON(201, category)
}
