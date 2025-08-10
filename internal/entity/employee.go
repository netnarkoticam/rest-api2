package employee

impot "time"

type Employee struct {
	ID string `json:"id" db:"id"`
	LastName string `json:"lastname" db:"lastname"`
	FirsName string `json:"firstname" db:"firstname"`
	MiddleName string `json:"middlename" db:"middlename"`
	Phone string `json:"phone" db:"phone"`
	Address string `json:"address" db:"address"`
	Department string `json:"department" db:"department"`
	HireDate time.Time `json:"hiredate" db:"hiredate"`
	FireDate time.Time `json:"firedate" db:"firedate"`
}
