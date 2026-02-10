package service

import (
	"context"
	"errors"

	"github.com/netnarkoticam/rest-api2.git/internal/entity"
	"github.com/netnarkoticam/rest-api2.git/internal/repo/pgdb"
)

type UserService struct {
	repo *pgdb.UserRepo
}

func NewUserService(repo *pgdb.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, user entity.User) (int, error) {
	if user.FirstName == "" {
		return 0, errors.New("first_name обязателен")
	}
	if user.PhoneNumber == "" {
		return 0, errors.New("phone_number обязателен")
	}

	return s.repo.RegisterUser(ctx, user)
}

func (s *UserService) GetUser(ctx context.Context, id int) (entity.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		return entity.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user entity.User) error {
	if user.ID == 0 {
		return errors.New("need id")
	}
	if user.FirstName == "" {
		return errors.New("need first name")
	}
	return s.repo.UpdateUser(ctx, user)
}

func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	return s.repo.DeleteUser(ctx, id)
}
