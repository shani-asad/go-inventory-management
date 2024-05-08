package database

import "time"

type Staff struct {
	Id			int
	PhoneNumber	string
	Password	string
	Name		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}