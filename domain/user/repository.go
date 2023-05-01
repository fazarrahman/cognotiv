package user

import (
	"context"

	ue "github.com/fazarrahman/cognotiv/domain/user/entity"
	"github.com/fazarrahman/cognotiv/error"
)

type UserRepository interface {
	GetUserByUsername(ctx context.Context, username string) (*ue.Users, *error.Error)
	GetUserByEmail(ctx context.Context, email string) (*ue.Users, *error.Error)
	GetRoleCodeByUsername(ctx context.Context, username string) (*string, *error.Error)
	InsertUser(ctx context.Context, user *ue.Users) *error.Error
}
