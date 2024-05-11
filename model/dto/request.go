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

type SearchSkuParams struct {
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Name string `json:"name"`
	Category string `json:"category"`
	Sku string `json:"sku"`
	Price string `json:"price"`
	InStock bool `json:"inStock"`
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
	CustomerID     string             `json:"customerId"`
	//ProductDetails []ProductDetailDTO `json:"productDetails"`
	Paid           int                `json:"paid"`
	Change         int                `json:"change"`
}
