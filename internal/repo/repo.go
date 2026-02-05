package repo

import (
	"context"

	"github.com/netnarkoticam/rest-api2.git/internal/entity"
)

type User interface {
	RegisterUser(ctx context.Context, user entity.User) (int, error)
}
