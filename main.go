package main

import "github.com/gin-gonic/gin"

func ShowStudents(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"name": "rafael",
	})
}

func main() {
	r := gin.Default()

	r.GET("/students", ShowStudents)

	r.Run()
}
