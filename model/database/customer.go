package database

import "time"

type Customer struct {
	Id			string
	PhoneNumber	string
	Name		string
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}