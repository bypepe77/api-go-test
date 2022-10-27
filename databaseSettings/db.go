package database

import (
	"fmt"
	"log"
	"os"

	models "test-api/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DatabaseConection() (gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&models.Movie{}, &models.Author{}, &models.User{})

	return *db, err
}
