package database

import "time"

type Product struct {
	ID          int
	Name        string
	SKU         string
	Category    string
	ImageURL    string
	Notes       string
	Price       float64
	Stock       int
	Location    string
	IsAvailable bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
