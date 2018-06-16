package models

import (
	"github.com/jinzhu/gorm"
)

// User represents a user of the application
type User struct {
	gorm.Model
	Address    string
	Messages   []Message
	LocationID uint
	Location   Location
}

// Message represents a message
type Message struct {
	gorm.Model
	Body   string
	UserID uint
}

// Location represents a location
type Location struct {
	gorm.Model
	Longitude float32
	Latitude  float32
	UserID    uint
}
