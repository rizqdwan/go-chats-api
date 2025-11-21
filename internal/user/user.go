package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID    		uuid.UUID `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	Email 		string    `json:"email" db:"email"`
	Password 	string 		`db:"password"`
}
