package dto

type RequestAuth struct {
	PhoneNumber    string `json:"phoneNumber"`
	Password string `json:"password"`
}

type RequestCreateStaff struct {
	PhoneNumber    string `json:"phoneNumber"`
	Password string `json:"password"`
	Name     string `json:"name"`
}