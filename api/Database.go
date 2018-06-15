package main

import (
	"log"
	"time"
	"twilio-playground/api/models"

	"github.com/jinzhu/gorm"
)

// CreateDatabaseConnection creates a connection the database
func CreateDatabaseConnection() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres", "postgresql://myapp:dbpass@localhost:15432/myapp")
	if err != nil {
		return
	}

	// Ping function checks the database connectivity
	err = db.DB().Ping()
	return
}

func addDatabase(db *gorm.DB, dbname string) error {
	// create database with dbname, won't do anything if db already exists
	db.Exec("CREATE DATABASE " + dbname)

	// connect to newly created DB (now has dbname param)
	connectionParams := "dbname=" + dbname + " user=docker password=docker sslmode=disable host=db"
	_, err := gorm.Open("postgres", connectionParams)
	if err != nil {
		return err
	}

	return nil
}

// IntializeDatabase creates tables for all of the models
func IntializeDatabase() (err error) {
	var db *gorm.DB
	// set up DB connection
	// then attempt to connect 10 times over 10 seconds
	connectionParams := "user=docker password=docker sslmode=disable host=db"
	for i := 0; i < 10; i++ {
		log.Printf("Trying to connect to database. Attempt: %v", i+1)
		db, err = gorm.Open("postgres", connectionParams)
		if err == nil {
			break
		}

		time.Sleep(1 * time.Second)
	}
	if err != nil {
		return
	}

	err = addDatabase(db, "twilio-playground")
	if err != nil {
		return
	}

	// create table if it does not exist
	if !db.HasTable(&models.Location{}) {
		db.CreateTable(&models.Location{})
	}
	if !db.HasTable(&models.Message{}) {
		db.CreateTable(&models.Message{})
	}
	if !db.HasTable(&models.User{}) {
		db.CreateTable(&models.User{})
	}

	return
}
