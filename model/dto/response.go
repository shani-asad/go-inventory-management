package dto

import "time"

type ToResponseGetCat struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type IssuedBy struct {
	Name string	`json:"name"`
	Email string	`json:"email"`
	CreatedAt string	`json:"createdAt"`
}

type CatDetail struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Race string `json:"race"`
	Sex string `json:"sex"`
	Description string `json:"description"`
	AgeInMonth int `json:"ageInMonth"`
	ImageUrls []string `json:"imageUrls"` // apa ini harusnya json string?
	HasMatched bool `json:"hasMatched"`
	CreatedAt string `json:"createdAt"`
}
type SingularResponseGetMatch struct {
	Id int		`json:"id"`
	IssuedBy IssuedBy `json:"issuedBy"`
	Message	string	`json:"message"`
	CreatedAt	string	`json:"createdAt"`
	MatchCatDetail CatDetail `json:"matchCatDetail"`
	UserCatDetail CatDetail `json:"userhCatDetail"`
}

type ResponseGetMatch struct {
	Data []SingularResponseGetMatch `json:"data"`
}
