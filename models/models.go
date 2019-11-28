package models

import "time"

// Flat : Model of Flat
type Flat struct {
	ID         int         `json:"id"`
	Name       string      `json:"name" binding:"required"`
	Price      int         `json:"price" binding:"required"`
	ResidentID interface{} `json:"residentID"`
	IsFree     bool        `json:"isFree" binding:"required"`
}

type FlatEdit struct {
	ID         int         `json:"id" binding:"required"`
	ResidentID interface{} `json:"residentID"`
}

// Resident : Model of Resident
type Resident struct {
	ID       int        `json:"id"`
	Name     string     `json:"name" binding:"required"`
	Contact  string     `json:"contact" binding:"required"`
	CheckIn  *time.Time `json:"checkIn" binding:"required"`
	CheckOut *time.Time `json:"checkOut" binding:"required"`
}
