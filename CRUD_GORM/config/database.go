package config

import (
	"fmt"
	"henry/helper"
	"os"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
)

var (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = os.Getenv("PASSWORD") // Replace "password" with the actual name of your environment variable
	dbName   = "test"
)

func DatabaseConnection() *gorm.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	helper.ErrorPanic(err)

	return db
}
