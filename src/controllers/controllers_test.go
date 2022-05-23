package controllers_test

import (
	"awesomeProject/src/controllers"
	"awesomeProject/src/models"
	"awesomeProject/src/repository"
	"awesomeProject/src/routes"
	"awesomeProject/src/server"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
var ctx *gin.Context

func TestListCategories(t *testing.T) {
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
