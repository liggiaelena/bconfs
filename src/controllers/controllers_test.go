package controllers_test

import (
	"awesomeProject/src/controllers"
	"awesomeProject/src/models"
	"awesomeProject/src/repository"
	"awesomeProject/src/routes"
	"awesomeProject/src/server"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSum(t *testing.T) {
	if controllers.Sum(2, 2) != 4 {
		t.Error("must be 4")
	}
}

var mockAdParams = []models.AdParams{{1, "name", "string", 1}}

var mockCategories = []models.Category{
	{1, "nameCategory", "description of the category", mockAdParams},
}
var mockCategoryNoExist = []models.Category{
	{0, "nameCategory", "description of the category", mockAdParams},
}
var ctx *gin.Context

func TestHandler_ListCategories_Status200(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetAllCategories").Return(mockCategories, nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, "/category", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)
	marshal, _ := json.Marshal(mockCategories)

	repo.AssertCalled(t, "GetAllCategories")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(marshal), w.Body.String())
}

func TestHandler_ListCategories_Status500(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetAllCategories").Return(mockCategories, errors.New("some error"))
	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, "/category", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "GetAllCategories")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandler_FindCategory_Status200(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryById", 1).Return(mockCategories, nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, `/category/1`, nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)
	marshal, _ := json.Marshal(mockCategories)

	repo.AssertCalled(t, "GetCategoryById", 1)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(marshal), w.Body.String())

}

func TestHandler_FindCategory_Status400(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryById").Return(mockCategories, nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, `/category/oi`, nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertNotCalled(t, "GetCategoryById")
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, `{"error":"id must be a number"}`, w.Body.String())

}

func TestHandler_FindCategory_Status500(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryById", 2).Return(mockCategories, errors.New("some error"))

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, `/category/2`, nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "GetCategoryById", 2)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandler_FindCategory_Status404(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryById", 2).Return(mockCategoryNoExist, nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, `/category/2`, nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "GetCategoryById", 2)
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestHandler_CreateCategory_Status200(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryByName", mock.Anything).Return(false)
	repo.On("PostCategory", mock.Anything).Return(mockCategories[0], nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)
	type Body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	aux := &Body{Name: "Anything", Description: "Some Description"}
	body, _ := json.Marshal(aux)

	req, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	marshal, _ := json.Marshal(mockCategories[0])

	repo.AssertCalled(t, "GetCategoryByName", mock.Anything)
	repo.AssertCalled(t, "PostCategory", mock.Anything)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Equal(t, string(marshal), w.Body.String())
}

func TestHandler_CreateCategory_Status400(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryByName", mock.Anything).Return(false)
	repo.On("PostCategory", mock.Anything).Return(mockCategories[0], nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)
	type Body struct {
		Name string `json:"name"`
	}
	aux := &Body{Name: "Anything"}
	body, _ := json.Marshal(aux)

	req, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertNotCalled(t, "GetCategoryByName", mock.Anything)
	repo.AssertNotCalled(t, "PostCategory", mock.Anything)
	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestHandler_CreateCategory_Status409(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryByName", mock.Anything).Return(true)
	repo.On("PostCategory", mock.Anything).Return(mockCategories[0], nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)
	type Body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	aux := &Body{Name: "Anything", Description: "Some Description"}
	body, _ := json.Marshal(aux)

	req, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "GetCategoryByName", mock.Anything)
	repo.AssertNotCalled(t, "PostCategory", mock.Anything)
	assert.Equal(t, http.StatusConflict, w.Code)
}

func TestHandler_CreateCategory_Status500(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("GetCategoryByName", mock.Anything).Return(false)
	repo.On("PostCategory", mock.Anything).Return(mockCategories[0], errors.New("some error"))

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)
	type Body struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}
	aux := &Body{Name: "Anything", Description: "Some Description"}
	body, _ := json.Marshal(aux)

	req, _ := http.NewRequest(http.MethodPost, "/category", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "GetCategoryByName", mock.Anything)
	repo.AssertCalled(t, "PostCategory", mock.Anything)
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestHandler_ListParams_Status200(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("ListParams").Return(mockAdParams, nil)

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, "/params", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)
	marshal, _ := json.Marshal(mockAdParams)

	repo.AssertCalled(t, "ListParams")
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, string(marshal), w.Body.String())
}

func TestHandler_ListParams_Status500(t *testing.T) {
	repo := &repository.RepositoryMock{}
	repo.On("ListParams").Return(mockAdParams, errors.New("some error"))

	server := server.NewServer()
	routes := routes.ConfigRoutes(server.Server, repo)

	req, _ := http.NewRequest(http.MethodGet, "/params", nil)
	w := httptest.NewRecorder()
	routes.ServeHTTP(w, req)

	repo.AssertCalled(t, "ListParams")
	assert.Equal(t, http.StatusInternalServerError, w.Code)
}
