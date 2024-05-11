package dto

type StaffLoginData struct {
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
	Message		int				`json:"message"`
	Data		StaffLoginData	`json:"data"`
}

type SkuData struct {
	Id			string	`json:"id"`
	Name		string	`json:"name"`
	Sku			string	`json:"sku"`
	Category	string	`json:"category"`
	ImageUrl	string	`json:"imageUrl"`
	Stock		int		`json:"stock"`
	Price		int		`json:"price"`
	Location	string	`json:"location"`
	CreatedAt	string	`json:"createdAt"`
}
type SearchSkuResponse struct {
	Message	string	`json:"message"`
	Data	SkuData	`json:"data"`
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