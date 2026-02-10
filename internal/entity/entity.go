package entity

type User struct {
	ID          int     `json:"id"`
	LastName    string  `json:"last_name"`
	FirstName   string  `json:"first_name"`
	MiddleName  string  `json:"middle_name"`
	PhoneNumber string  `json:"phone_number"`
	Address     string  `json:"address"`
	Department  string  `json:"department"`
	HireDate    string  `json:"hire_date"`
	FireDate    *string `json:"fire_date"`
}
