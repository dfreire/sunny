package model

func CustomerRoles() []string {
	return []string{
		"sommelier",
		"restaurant",
		"wine_distribution",
		"wine_shop",
		"wine_lover",
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
	Id        string `json:"id",        dbx:"id"`
	Email     string `json:"email",     dbx:"email"`
	Role      string `json:"role",      dbx:"role"`
	CreatedAt string `json:"createdAt", dbx:"createdAt"`
}

type WineComment struct {
	Id         string `json:"id",         dbx:"id"`
	CustomerId string `json:"customerId", dbx:"customerId"`
	WineId     string `json:"wineId",     dbx:"wineId"`
	WineYear   int    `json:"wineYear",   dbx:"wineYear"`
	CreatedAt  string `json:"createdAt",  dbx:"createdAt"`
	UpdatedAt  string `json:"updatedAt",  dbx:"updatedAt"`
	Comment    string `json:"comment",    dbx:"comment"`
}
