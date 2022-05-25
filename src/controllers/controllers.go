package controllers

import (
	"awesomeProject/src/models"
	"awesomeProject/src/repository"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func Sum(num1 int, num2 int) int {
	x := num1 + num2
	return x
}

type (
	IHandler interface {
		ListCategories(c *gin.Context)
		FindCategory(c *gin.Context)
		CreateCategory(c *gin.Context)
		ListParams(c *gin.Context)
	}

	Handler struct {
		Repo repository.Repository
	}
)

func NewHandler(repo repository.Repository) IHandler {
	return &Handler{
		Repo: repo,
	}
}

func (h *Handler) ListCategories(c *gin.Context) {
	category, err := h.Repo.GetAllCategories()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, category)
}

func (h *Handler) FindCategory(c *gin.Context) {
	param := c.Param("id")

	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(400, gin.H{"error": "id must be a number"})
		return
	}

	category, err := h.Repo.GetCategoryById(id)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	if category[0].ID == 0 {
		c.JSON(404, gin.H{
			"error": "this category does not exist",
		})
		return
	}
	c.JSON(200, category)
}

func (h *Handler) CreateCategory(c *gin.Context) {
	var category models.Category

	err := c.ShouldBindJSON(&category) //how get the body?
	fmt.Println(category)
	fmt.Println(err)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	exist := h.Repo.GetCategoryByName(category.Name)
	fmt.Println(category)
	if exist {
		c.JSON(409, gin.H{
			"error": "This category already exist",
		})
		return
	}

	categories, err := h.Repo.PostCategory(category)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "cannot create category: " + err.Error(),
		})
		return
	}
	c.JSON(201, categories[0])
}

func (h *Handler) ListParams(c *gin.Context) {
	params, err := h.Repo.ListParams()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, params)
}
