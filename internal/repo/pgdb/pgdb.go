package pgdb

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/netnarkoticam/rest-api2.git/internal/entity"
)

type UserRepo struct{
	DB *sql.DB
}

func NewUserRepo (db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) RegisterUser (ctx context.Context, user entity.User) (int, error){
	query := `INSERT INTO employees (last_name, first_name, middle_name, phone_number, address, department, hire_date, fire_date) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
	var id int
	err := r.DB.QueryRowContext(ctx, query, user.LastName, user.FirstName, user.MiddleName, user.PhoneNumber, user.Address, user.Department, user.HireDate, user.FireDate).Scan(&id)
	return id, err
}