package database

import (
	"github.com/rafawildfire42/go-gin-api-study/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDatabase() {
	connectionString := "host=localhost user=root password=root dbname=root port=5423 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		panic("Erro ao conectar no banco de dados.")
	}
	DB.AutoMigrate(&models.Student{})
}
