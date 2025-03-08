package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateUser(ctx context.Context, conn *pgxpool.Pool, user *User) (User, error) {
	query := `
		INSERT INTO users (name, surname, email)
		VALUES ($1, $2, $3)
		RETURNING id, name, surname, email, created_at
	`
	var createdUser User
	err := conn.QueryRow(ctx, query, user.Name, user.Surname, user.Email).Scan(
		&createdUser.ID,
		&createdUser.Name,
		&createdUser.Surname,
		&createdUser.Email,
		&createdUser.CreatedAt,
	)
	if err != nil {
		return User{}, err
	}
	return createdUser, nil
}
