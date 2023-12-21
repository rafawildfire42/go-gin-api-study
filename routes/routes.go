package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafawildfire42/go-gin-api-study/controllers"
)

func HandleRequets() {
	r := gin.Default()

	r.GET("/students", controllers.ListStudents)
	r.GET("/students/:id", controllers.GetStudent)
	r.GET("/students/cpf/:cpf", controllers.GetStudentByCPF)

	r.POST("/students", controllers.CreateStudent)

	r.PATCH("/students/:id", controllers.UpdateStudent)

	r.DELETE("/students/:id", controllers.DeleteStudent)

	r.GET("/test-me", controllers.EndpointToTest)

	r.Run()
}
