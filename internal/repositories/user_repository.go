package repositories

import (
	"context"

	"github.com/rizqdwan/go-chats-api/internal/user"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u user.User) (*user.User, error)
	GetUserByEmail(ctx context.Context, email string) (*user.User, error)
}
