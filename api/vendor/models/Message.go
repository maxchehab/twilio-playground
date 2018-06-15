package models

import "time"

// Message represents a message
type Message struct {
	ID          int `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Body        string
	CreatedDate time.Time `sql:"DEFAULT:current_timestamp"`
}
