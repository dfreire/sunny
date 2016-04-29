package model

// type Customer struct {
// 	Id             string      `json:"id,omitempty",dbx:"id"`
// 	Name           null.String `json:"name,omitempty",dbx:"name"`
// 	Email          string      `json:"email,omitempty",dbx:"email"`
// 	RoleId         string      `json:"roleId,omitempty",dbx:"roleId"`
// 	CreatedAt      string      `json:"createdAt,omitempty",dbx:"createdAt"`
// 	SignupOriginId string      `json:"signupOriginId,omitempty",dbx:"signupOriginId"`
// 	InMailingList  bool        `json:"inMailingList",dbx:"inMailingList"`
// }

type WineComment struct {
	Id         string `json:"id,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	WineId     string `json:"wineId,omitempty"`
	WineYear   int    `json:"wineYear,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
	Comment    string `json:"comment,omitempty"`
}
