package model

type Customer struct {
	Id             string `json:"id,omitempty"`
	Email          string `json:"email,omitempty"`
	RoleId         string `json:"roleId,omitempty"`
	CreatedAt      string `json:"createdAt,omitempty"`
	SignupOriginId string `json:"signupOriginId,omitempty"`
	InMailingList  bool   `json:"inMailingList,omitempty"`
}

type WineComment struct {
	Id         string `json:"id,omitempty"`
	CustomerId string `json:"customerId,omitempty"`
	WineId     string `json:"wineId,omitempty"`
	WineYear   int    `json:"wineYear,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	UpdatedAt  string `json:"updatedAt,omitempty"`
	Comment    string `json:"comment,omitempty"`
}
