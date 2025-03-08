package users

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

// CreateUserHandler godoc
// @Summary Create a new user
// @Description Create a new user with the provided information
// @Tags users
// @Accept json
// @Produce json
// @Param user body UserRequest true "User information"
// @Success 201 {object} UserResponse "User created successfully"
// @Failure 400 {object} ErrorResponse "Invalid request"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users [post]
func CreateUserHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		if name == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Name is required"})
			return
		}
		surname := r.FormValue("surname")
		if surname == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Surname is required"})
			return
		}
		email := r.FormValue("email")
		if email == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Email is required"})
			return
		}
		user := User{
			Name:    name,
			Surname: surname,
			Email:   email,
		}

		newUser, err := CreateUser(context.Background(), db, &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newUser)
	}
}

// ReadManyUsersHandler godoc
// @Summary Get all users
// @Description Retrieve a list of all users
// @Tags users
// @Produce json
// @Param limit query int false "Limit the number of results" default(100) minimum(1) maximum(1000)
// @Param offset query int false "Offset for pagination" default(0) minimum(0)
// @Success 200 {array} UserResponse "List of users"
// @Failure 500 {object} ErrorResponse "Server error"
// @Router /users [get]
func ReadManyUsersHandler(db *pgxpool.Pool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := ReadManyUsers(context.Background(), db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(users)
	}
}

// ReadUserHandler godoc
// @Summary Get a user by ID
// @Description Retrieve a user by their unique identifier
// @Tags users

