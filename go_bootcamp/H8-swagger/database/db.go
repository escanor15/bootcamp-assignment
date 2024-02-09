package database

import (
	"fmt"
	"go_bootcamp/H8-swagger/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "04091997"
	dbPort   = "5432"
	dbname   = "learning-gorm"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connectiong to database :", err)

	}

	db.Debug().AutoMigrate(models.Car{})
}

func GetDB() *gorm.DB {

	dsn := "user=postgres dbname=learning-gorm password=04091997 host=localhost port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Error connecting to the database:", err)
		return nil
	}
	return db
}
