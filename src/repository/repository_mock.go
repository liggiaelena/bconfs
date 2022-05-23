package repository

import (
	"awesomeProject/src/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetAllCategories() ([]models.Category, error) {
	return nil, nil

}
func (m *RepositoryMock) GetCategoryById(id int) ([]models.Category, error) {
	return nil, nil

}
func (m *RepositoryMock) GetCategoryByName(name string) bool {
	return false

}
func (m *RepositoryMock) PostCategory(category models.Category) ([]models.Category, error) {
	return nil, nil

}
func (m *RepositoryMock) ListParams() ([]models.AdParams, error) {
	return nil, nil

}
