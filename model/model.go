package model

func CustomerRoles() []string {
	return []string{
		"sommelier",
		"restaurant",
		"Wine_distribution",
		"Wine_shop",
		"Wine_lover",
		"other",
	}
}

func IsCustomerRole(role string) bool {
	for _, r := range CustomerRoles() {
		if r == role {
			return true
		}
	}
	return false
}

type Customer struct {
	Id        string `json:"id,omitempty"`
	Email     string `json:"email,omitempty"`
	Role      string `json:"role,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
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
