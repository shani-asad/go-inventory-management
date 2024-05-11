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
