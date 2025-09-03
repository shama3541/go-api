package store

import (
	"context"
	"database/sql"
)

type UserStore struct {
	db *sql.DB
}

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdateAt  string `json:"updated_at"`
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `INSERT INTO users (username, password, created_at, updated_at)
			  VALUES($1, $2, NOW(), NOW()) RETURNING id, created_at, updated_at`

	err := s.db.QueryRowContext(ctx, query, user.Username, user.Password).Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdateAt,
	)
	if err != nil {
		return err
	}
	return nil

}
