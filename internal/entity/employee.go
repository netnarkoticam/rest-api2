package employee

import "time"

type Employee struct {
	ID         string    `json:"id" db:"id"`
	LastName   string    `json:"lastname" db:"last_name"`
	FirsName   string    `json:"firstname" db:"first_name"`
	MiddleName string    `json:"middlename" db:"middle_name"`
	Phone      string    `json:"phone" db:"phone"`
	Address    string    `json:"address" db:"address"`
	Department string    `json:"department" db:"department"`
	HireDate   time.Time `json:"hiredate" db:"hire_date"`
	FireDate   time.Time `json:"firedate" db:"fire_date"`
}
