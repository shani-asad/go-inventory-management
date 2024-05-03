package dto

type RequestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestCreateCat struct {
	Name 					string `json:"name"`
	Race 					string `json:"race"`
	Sex 					string `json:"sex"`
	AgeInMonth 		int `json:"ageInMonth"`
	Description 	string `json:"description"`
	ImageUrls 		[]string `json:"imageUrls"`
}
