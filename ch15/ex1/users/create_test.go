package users

import (
	"context"
	"testing"
)

func TestCreateUser(t *testing.T) {
	user := User{
		Name:  "John Doe",
		Email: "john@example.com",
	}

	_, err := CreateUser(context.Background(), testDB, &user)
	if err != nil {
		t.Fatalf("Failed to create user: %v", err)
	}

	users, err := ReadManyUsers(context.Background(), testDB)
	if err != nil {
		t.Fatalf("Failed to read users: %v", err)
	}

	if len(users) < 1 {
		t.Fatalf("Expected at least 1 user, got %d", len(users))
	}
}
