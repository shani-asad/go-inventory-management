package database

import "time"

type Cat struct {
	Id        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
