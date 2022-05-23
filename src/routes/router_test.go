package routes_test

import (
	"awesomeProject/src/routes"
	"fmt"
	"github.com/gin-gonic/gin"
	"testing"
)

func TestRoutes(t *testing.T) {
	//var app *gin.Engine
	app1 := &gin.Engine{RedirectTrailingSlash: false}

	mux := routes.ConfigRoutes(app1)

	switch v := interface{}(mux).(type) {
	case *gin.Engine:
		fmt.Println("oi")
	default:
		t.Error(fmt.Sprintf("type is not gin.Engine, but is %T", v))
	}
}
