package pgdb

import (
	"context"
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq" // иначе не работает
	"github.com/netnarkoticam/rest-api2.git/internal/entity"
)

type UserRepo struct {
	DB *sql.DB
}

func NewUserRepo(db *sql.DB) *UserRepo {
	return &UserRepo{DB: db}
}

func (r *UserRepo) RegisterUser(ctx context.Context, user entity.User) (int, error) {
	log.Printf("USER DATA: %+v", user)
	const query = `INSERT INTO employees (first_name, last_name, middle_name, phone_number, address, department) 
    VALUES ($1, $2, COALESCE($3, ''), $4, COALESCE($5, ''), COALESCE($6, '')) 
    RETURNING id`
	log.Printf("EXECUTING QUERY with params: first_name=%s, last_name=%s, phone=%s",
		user.FirstName, user.LastName, user.PhoneNumber)
	var id int
	err := r.DB.QueryRowContext(ctx, query,
		user.LastName,
		user.FirstName,
		user.MiddleName,
		user.PhoneNumber,
		user.Address,
		user.Department,
	).Scan(&id)
	if err != nil {
		log.Printf("DB ERROR: %v", err)
		return 0, errors.New("register failed")
	}
	log.Printf("User created with ID: %d", id)
	return id, err
}

func (r *UserRepo) GetUserByID(ctx context.Context, id int) (entity.User, error) {
	log.Printf("GetUserId : %d", id)
	var user entity.User
	query := `SELECT id, last_name, first_name, middle_name, phone_number, address, department
	          FROM employees WHERE id = $1`
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&user.ID, &user.LastName, &user.FirstName, &user.MiddleName, &user.PhoneNumber, &user.Address, &user.Department)
	if err != nil {
		log.Printf("GetUser ERROR ID=%d: %v", id, err)

		return user, errors.New("user not found")
	}
	log.Printf("SUCCES ID = %d", id)
	return user, nil
}

func (r *UserRepo) UpdateUser(ctx context.Context, user entity.User) error {
	query := `UPDATE employees
	          SET first_name = $1, last_name = $2, phone_number = $3
			  WHERE id = $4`
	result, err := r.DB.ExecContext(ctx, query, user.FirstName, user.LastName, user.PhoneNumber, user.ID)
	if err != nil {
		return errors.New("failed update")
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return errors.New("failde update")
	}
	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}

func (r *UserRepo) DeleteUser(ctx context.Context, id int) error {
	const query = `DELETE FROM employees WHERE id = $1`

	res, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return errors.New("deleting has been failed")
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return errors.New("rows check failed")
	}

	if rows == 0 {
		return errors.New("user not found")
	}
	return nil
}
