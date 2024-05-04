package database

import "time"

type Match struct {
	Id			int	
	MatchCatId	int	// cat ID
	UserCatId	int // cat ID
	Message		string // not null, minLength: 5, maxLength: 120
	CreatedAt	time.Time
	UpdatedAt	time.Time
}
