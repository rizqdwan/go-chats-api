package user

import "github.com/google/uuid"

type CreateUserRequest struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username" db:"username"`
	Email    string    `json:"email" db:"email"`
}
