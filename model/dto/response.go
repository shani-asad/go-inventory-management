package dto

import "time"

type StaffData struct {
	UserId      string `json:"userId"`
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
	AccessToken string `json:"accessToken"`
}
type RegistrationResponse struct {
	Message int       `json:"message"`
	Data    StaffData `json:"data"`
}

type RegisterCustomerResponse struct {
	Message string `json:"message"`
	//Data    CustomerDTO `json:"data"`
}

type CheckoutProductResponse struct {
	Message string `json:"message"`
}

type ResponseCreateProduct struct {
	Message string `json:"message"`
	Data    struct {
		ID        string    `json:"id"`
		CreatedAt time.Time `json:"createdAt"`
	} `json:"data"`
}

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	SKU         string    `json:"sku"`
	Category    string    `json:"category"`
	ImageURL    string    `json:"imageUrl"`
	Stock       int       `json:"stock"`
	Notes       string    `json:"notes"`
	Price       float64   `json:"price"`
	Location    string    `json:"location"`
	IsAvailable bool      `json:"isAvailable"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ResponseGetProduct struct {
	Message string    `json:"message"`
	Data    []Product `json:"data"`
}
