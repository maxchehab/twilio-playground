package models

import "time"

// User represents a user of the application
type User struct {
	ID          int    `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Name        string `sql:"size:255;unique;index"`
	Address     string
	Messages    []Message
	Location    Location
	CreatedDate time.Time `sql:"DEFAULT:current_timestamp"`
}
