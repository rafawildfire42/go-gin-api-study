package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/rafawildfire42/go-gin-api-study/controllers"
	"github.com/rafawildfire42/go-gin-api-study/database"
	"github.com/stretchr/testify/assert"
)

func SetupTestRoutes() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	routes := gin.Default()
	return routes
}

func TestMyRoute(t *testing.T) {
	r := SetupTestRoutes()
	r.GET("/test-me", controllers.EndpointToTest)
	request, _ := http.NewRequest("GET", "/test-me", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code, "Should be equal!")

	responseMock := `{"message":"Deu tudo certo!"}`
	bodyResponse, _ := (io.ReadAll(response.Body))

	assert.Equal(t, responseMock, string(bodyResponse))

}

func TestListStudents(t *testing.T) {
	database.ConnectDatabase()
	r := SetupTestRoutes()
	r.GET("/students", controllers.ListStudents)
	request, _ := http.NewRequest("GET", "/students", nil)
	response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	assert.Equal(t, http.StatusOK, response.Code)
}
