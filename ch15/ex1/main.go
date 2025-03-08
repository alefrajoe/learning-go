package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"webserver/users"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// @title User API
// @version 1.0
// @description A simple user management API
// @host localhost:8080
// @BasePath /
// @schemes http
// @produce json
// @consumes json
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Connect to database
	conn, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// Spin up server
	mux := http.NewServeMux()
	mux.HandleFunc("POST /users", users.CreateUserHandler(conn))
	mux.HandleFunc("GET /users", users.ReadManyUsersHandler(conn))

	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
