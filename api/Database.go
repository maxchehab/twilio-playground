package main

import (
	"log"
	"models"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var database *gorm.DB

// IntializeDatabase creates tables for all of the models
func IntializeDatabase() (err error) {
	// set up DB connection
	// then attempt to connect 10 times over 10 seconds
	connectionParams := "user=docker password=docker sslmode=disable host=db"
	connectionWithDatabaseParams := "user=docker password=docker sslmode=disable host=db dbname=twilio_playground"

	for i := 0; i < 30; i++ {
		log.Printf("trying to connect to database, attempt: %v", i+1)
		database, err = gorm.Open("postgres", connectionParams)
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}

		log.Printf("creating database '%v'", "twilio_playground")
		database.Exec("CREATE DATABASE " + "twilio_playground")

		database, err = gorm.Open("postgres", connectionWithDatabaseParams)
		if err == nil {
			break
		}
	}
	if err != nil {
		return
	}

	database.AutoMigrate(&models.Location{})
	database.AutoMigrate(&models.Message{})
	database.AutoMigrate(&models.User{})

	return
}

// SaveToDatabase saves a message and returns a User
func SaveToDatabase(address string, body string) (user models.User, err error) {
	user = models.User{Address: address}
	if result := database.Where(&models.User{Address: address}).FirstOrCreate(&user); result.Error != nil {
		err = result.Error
		return
	}

	if result := database.Model(&user).Association("Messages").Append(models.Message{Body: body}); result.Error != nil {
		err = result.Error
		return
	}
	return
}
