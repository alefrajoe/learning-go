package users

import "time"

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
