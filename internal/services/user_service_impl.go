package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rizqdwan/go-chats-api/internal/repositories"
	"github.com/rizqdwan/go-chats-api/internal/user"
	"github.com/rizqdwan/go-chats-api/util"
)

type userServiceImpl struct {
	repo    repositories.UserRepository
	jwt     JWTService
	timeout time.Duration
}

func NewUserService(repo repositories.UserRepository, jwt JWTService) UserService {
	return &userServiceImpl{
		repo:    repo,
		jwt:     jwt,
		timeout: time.Duration(2) * time.Second,
	}
}

func (s *userServiceImpl) CreateUser(ctx context.Context, req *user.CreateUserRequest) (*user.CreateUserResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// Check if user exists
	existingUser, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// Create user
	newUser := &user.User{
		ID:       uuid.New(),
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	createdUser, err := s.repo.CreateUser(ctx, newUser)
	if err != nil {
		return nil, err
	}

	return &user.CreateUserResponse{
		ID:       createdUser.ID,
		Username: createdUser.Username,
		Email:    createdUser.Email,
	}, nil
}

func (s *userServiceImpl) SignIn(ctx context.Context, req *user.LoginUserRequest) (*user.LoginUserResponse, string, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	// Get user
	u, err := s.repo.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, "", err
	}
	if u == nil {
		return nil, "", errors.New("invalid credentials")
	}

err = util.ComparePassword(req.Password, u.Password)
	if err != nil {
		return &user.LoginUserResponse{}, err
	}

	// Generate token
	token, err := s.jwt.GenerateToken(u.ID.String(), u.Email)
	if err != nil {
		return nil, "", err
	}

	return &user.LoginUserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
	}, token, nil
}