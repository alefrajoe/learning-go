package users

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

// Add a package-level variable
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	// Set the package-level variable
	var err2 error
	testDB, err2 = pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err2 != nil {
		log.Fatalf("Failed to connect to database: %v", err2)
	}
	defer testDB.Close()

	os.Exit(m.Run())
}

func TestReadManyUsersHandler(t *testing.T) {
	users, err := ReadManyUsers(context.Background(), testDB)
	if err != nil {
		t.Fatalf("Failed to read users: %v", err)
	}

	if len(users) == 0 {
		t.Fatalf("No users found")
	}
}
