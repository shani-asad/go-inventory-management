package dto

type RequestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RequestCreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type RequestCreateMatch struct {
	MatchCatId	int `json:"matchCatId"`
	UserCatId	int `json:"userCatId"`
	Message		string `json:"message"`
}
type RequestCreateCat struct {
	Name 					string `json:"name"`
	UserId 				int `json:"userId"`
	Race 					string `json:"race"`
	Sex 					string `json:"sex"`
	AgeInMonth 		int `json:"ageInMonth"`
	Description 	string `json:"description"`
	ImageUrls 		[]string `json:"imageUrls"`
}

type RequestGetCat struct {
	Id string `json:"id"`
	Limit int `json:"limit"`
	Offset int `json:"offset"`
	Race string `json:"race"`
	Sex string `json:"sex"`
	HasMatched bool `json:"hasMatched"`
	AgeInMonth string `json:"ageInMonth"`
	Owned bool `json:"owned"`
	Search string `json:"search"`
}

type RequestApproveMatch struct {
	MatchId int `json:"matchId"`
}
