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