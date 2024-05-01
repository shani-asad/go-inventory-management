package database

import "time"

type Cat struct {
	Id        	int
	Name      	string
	Race		string
	Sex			string
	AgeInMonth	int
	Description	string // not null, minLength 1, maxLength 200
	CreatedAt	time.Time
	UpdatedAt	time.Time
	ImageUrls	[]string // not null, minItems: 1, items: not null, should be url
}