package database

import "time"

type Match struct {
	Id			int
	IssuedBy	int // user ID
	MatchCatId	int	// cat ID
	UserCatId	int // cat ID
	Message		int // not null, minLength: 5, maxLength: 120
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
