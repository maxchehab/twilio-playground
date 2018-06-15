package models

import "time"

// Location represents a Location
type Location struct {
	Longitude   float32
	Latitude    float32
	CreatedDate time.Time `sql:"DEFAULT:current_timestamp"`
}
