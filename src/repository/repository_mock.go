package repository

import (
	"awesomeProject/src/models"
	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	mock.Mock
}

func (m *RepositoryMock) GetAllCategories() ([]models.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Category), args.Error(1)
}

func (m *RepositoryMock) GetCategoryById(id int) ([]models.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Category), args.Error(1)
}

func (m *RepositoryMock) GetCategoryByName(name string) bool {
	args := m.Called()
	result := args.Get(0)
	return result.(bool)
}

func (m *RepositoryMock) PostCategory(category models.Category) ([]models.Category, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.Category), args.Error(1)
}

func (m *RepositoryMock) ListParams() ([]models.AdParams, error) {
	args := m.Called()
	result := args.Get(0)
	return result.([]models.AdParams), args.Error(1)
}
