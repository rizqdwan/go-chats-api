package services

import (
	"context"

	"github.com/rizqdwan/go-chats-api/internal/user"
)

type UserService interface {
	CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error)
	SignIn(ctx context.Context, req *user.LoginUserRequest) (*user.LoginUserResponse, string, error)
}