package dto

type StaffData struct {
	UserId			string 		`json:"userId"`
	PhoneNumber		string 		`json:"phoneNumber"`
	Name			string 		`json:"name"`
	AccessToken		string 		`json:"accessToken"`
}

type CustomerDTO struct {
	Id        string     `json:"id"`
	Name      string     `json:"name"`
	PhoneNumber string     `json:"phoneNumber"`
}
type RegistrationResponse struct {
	Message			int			`json:"message"`
	Data			StaffData	`json:"data"`
}

type RegisterCustomerResponse struct {
	Message string      `json:"message"`
	Data    CustomerDTO `json:"data"`
}

type SearchCustomersResponse struct {
	Message string         `json:"message"`
	Data    []CustomerDTO `json:"data"`
}


type CheckoutProductResponse struct {
	Message string `json:"message"`
}