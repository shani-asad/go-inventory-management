package dto

import "time"

type RequestAuth struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

type RequestCreateStaff struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
	Name        string `json:"name"`
}

type SearchSkuParams struct {
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	Sku            string `json:"sku"`
	Price          string `json:"price"`
	InStock        bool   `json:"inStock"`
	IsInstockValid bool
}
type RegisterCustomerRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type SearchCustomersRequest struct {
	PhoneNumber string `json:"phoneNumber"`
	Name        string `json:"name"`
}

type CheckoutProductRequest struct {
	CustomerID string `json:"customerId"`
	//ProductDetails []ProductDetailDTO `json:"productDetails"`
	Paid   int `json:"paid"`
	Change int `json:"change"`
}

type RequestGetProduct struct {
	ID          string `form:"id"`
	Limit       int    `form:"limit"`
	Offset      int    `form:"offset"`
	Name        string `form:"name"`
	IsAvailable string `form:"isAvailable"`
	Category    string `form:"category"`
	SKU         string `form:"sku"`
	Price       string `form:"price"`
	Instock     string `form:"inStock"`
	CreatedAt   string `form:"createdAt"`
}

type RequestUpsertProduct struct {
	ID          int       `json:"-"`
	Name        string    `json:"name" validate:"required,min=1,max=30"`
	SKU         string    `json:"sku" validate:"required,min=1,max=30"`
	Category    string    `json:"category" validate:"required,categoryEnum"`
	ImageURL    string    `json:"imageUrl" validate:"required,url"`
	Notes       string    `json:"notes" validate:"required,min=1,max=200"`
	Price       float64   `json:"price" validate:"required,min=1"`
	Stock       int       `json:"stock" validate:"required,min=1,max=100000"`
	Location    string    `json:"location" validate:"required,min=1,max=200"`
	IsAvailable *bool     `json:"isAvailable" validate:"required"`
	CreatedAt   time.Time `json:"createdAt"`
}
