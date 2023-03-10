package database

import (
	"log"

	"github.com/thifelipesilva/go_gin_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConecctionWithDB() {
	addressConecction := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(addressConecction))
	if err != nil {
		log.Panic(err)
	}

	DB.AutoMigrate(&models.Student{}) //endereco de memoria da instancia da struct
}
