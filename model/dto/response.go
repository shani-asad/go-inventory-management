package dto

type StaffData struct {
	UserId			string 		`json:"userId"`
	PhoneNumber		string 		`json:"phoneNumber"`
	Name			string 		`json:"name"`
	AccessToken		string 		`json:"accessToken"`
}
type RegistrationResponse struct {
	Message			int			`json:"message"`
	Data			StaffData	`json:"data"`
}

type RegisterCustomerResponse struct {
	Message string      `json:"message"`
	//Data    CustomerDTO `json:"data"`
}

type CheckoutProductResponse struct {
	Message string `json:"message"`
}