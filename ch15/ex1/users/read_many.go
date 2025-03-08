package users

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ReadManyUsers(ctx context.Context, conn *pgxpool.Pool) ([]User, error) {
	query := `SELECT id, name, surname, email, created_at FROM users`

	rows, err := conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []User{}
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Name, &user.Surname, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
