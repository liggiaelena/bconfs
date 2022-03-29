package controllers

import "github.com/gin-gonic/gin"

func ListCategories(c *gin.Context) {
	c.JSON(
		200,
		gin.H{
			"valuek": "ola",
		},
	)
}
