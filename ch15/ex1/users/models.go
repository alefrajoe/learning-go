package users

// UserRequest represents the request body for creating a user
// @Description User creation request
type UserRequest struct {
	// The user's name
	// required: true
	// example: John Doe
	Name string `json:"name"`

	// The user's email address
	// required: true
	// example: john@example.com
	Email string `json:"email"`

	// The user's age
	// min: 0
	// max: 120
	// example: 30
	Age int `json:"age,omitempty"`
}

// UserResponse represents the response for user operations
// @Description User response
type UserResponse struct {
	// The user's unique identifier
	// example: 123e4567-e89b-12d3-a456-426614174000
	ID string `json:"id"`

	// The user's name
	// example: John Doe
	Name string `json:"name"`

	// The user's email address
	// example: john@example.com
	Email string `json:"email"`

	// The user's age
	// example: 30
	// required: false
	Age int `json:"age,omitempty"`
}

// ErrorResponse represents an error response
// @Description Error response
type ErrorResponse struct {
	// The error message
	// example: Invalid request parameters
	Message string `json:"message"`
}
