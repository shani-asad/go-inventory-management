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