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
