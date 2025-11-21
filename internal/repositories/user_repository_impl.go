package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/rizqdwan/go-chats-api/internal/user"
)

type UserRepositoryImpl struct {
	db DBTX
}

func NewUserRepository(db DBTX) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (r *UserRepositoryImpl) CreateUser(ctx context.Context, u user.User) (*user.User, error) {
	u.ID = uuid.New()

	query := `
		INSERT INTO users (id, username, email, password)
		VALUES ($1, $2, $3, $4)
		RETURNING id
	`

	err := r.db.QueryRowContext(ctx, query, u.ID, u.Username, u.Email, u.Password).Scan(&u.ID)
	if err != nil {
		return nil, err
	}

	return &u, nil
}

func (r *UserRepositoryImpl) GetUserByEmail(ctx context.Context, email string) (*user.User, error) {
	var u user.User

	query := `
		SELECT id, username, email, password
		FROM users
		WHERE email = $1
	`

	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}