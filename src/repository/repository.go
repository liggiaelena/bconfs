package repository

import (
	"awesomeProject/src/database"
	"awesomeProject/src/models"
	"gorm.io/gorm"
)

type (
	Repository interface {
		GetAllCategories() ([]models.Category, error)
		GetCategoryById(id int) ([]models.Category, error)
		GetCategoryByName(name string) bool
		PostCategory(category models.Category) ([]models.Category, error)
		ListParams() ([]models.AdParams, error)
	}

	repository struct {
		conn *gorm.DB
	}
)

func DefaultRepository() Repository {
	return &repository{
		conn: database.GetDatabase(),
	}
}

func (r *repository) GetAllCategories() ([]models.Category, error) {
	//db := database.GetDatabase()
	var category []models.Category
	err := r.conn.Preload("AdParams").Find(&category).Error
	if err != nil {
		//jogar um erro
		return nil, err
	}
	return category, nil
}

func (r *repository) GetCategoryById(id int) ([]models.Category, error) {
	//db := database.GetDatabase()
	var category models.Category
	err := r.conn.Preload("AdParams").Find(&category, id).Error

	if err != nil {
		//jogar um erro
		return nil, err
	}
	return []models.Category{category}, nil
}

func (r *repository) GetCategoryByName(name string) bool {
	//db := database.GetDatabase()
	var category models.Category
	err := r.conn.Find(&category, "name = ?", category.Name).Error

	if err != nil {
		return true
	}
	return false
}

func (r *repository) PostCategory(category models.Category) ([]models.Category, error) {
	//db := database.GetDatabase()

	err := r.conn.Create(&category).Error

	if err != nil {
		return nil, err
	}
	return []models.Category{category}, nil
}

func (r *repository) ListParams() ([]models.AdParams, error) {
	//db := database.GetDatabase()
	var params []models.AdParams
	err := r.conn.Find(&params).Error
	if err != nil {
		return nil, err
	}
	return params, nil
}
