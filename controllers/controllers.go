package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rafawildfire42/go-gin-api-study/database"
	"github.com/rafawildfire42/go-gin-api-study/models"
	"gopkg.in/validator.v2"
)

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	if err := validator.Validate(student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&student)

	c.JSON(http.StatusOK, student)
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Model(&student).UpdateColumns(student)

	c.JSON(http.StatusOK, gin.H{
		"message": "Student update.",
	})

}

func ListStudents(c *gin.Context) {
	var students []models.Student
	database.DB.Find(&students)
	c.JSON(http.StatusOK, students)
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.First(&student, id)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Params.ByName("id")
	database.DB.Delete(&student, id)

	c.JSON(http.StatusOK, student)
}

func GetStudentByCPF(c *gin.Context) {
	var student models.Student
	cpf := c.Params.ByName("cpf")
	database.DB.First(&student, "cpf = ?", cpf)

	if student.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Student not found",
		})
		return
	}
	c.JSON(http.StatusOK, student)
}

func EndpointToTest(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Deu tudo certo!",
	})

}
