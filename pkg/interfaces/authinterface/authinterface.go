package authinterface

import "github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/authmodels"

// AuthInterface defines the methods for user-related operations.
type AuthInterface interface {
	// CreateUser creates a new user in the database.
	// Parameters:
	//   - r: CreateUserRequest containing username, email, and password
	// Returns:
	//   - error: nil if successful, otherwise contains the error message
	CreateUser(r *authmodels.CreateUserRequest) error

	// GetUserByEmail retrieves user information using their email address.
	// Parameters:
	//   - email: string containing the user's email address
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByEmail(email string) (*authmodels.User, error)

	// GetUserByID retrieves user information using their unique identifier.
	// Parameters:
	//   - id: integer representing the user's unique ID in the database
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByID(id int) (*authmodels.User, error)

	// GetUserByUsername retrieves user information using their username.
	// Parameters:
	//   - username: string containing the user's chosen username
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByUsername(username string) (*authmodels.User, error)

	// CheckUsernameExists verifies if a username is already taken in the database.
	// Parameters:
	//   - username: string containing the username to check
	// Returns:
	//   - bool:
	//     * true: username is already registered by another user and cannot be used
	//     * false: username is available and can be used for registration
	//   - error: nil if check successful, otherwise contains database error details
	CheckUsernameExists(username string) (bool, error)

	// UpdatePassword changes a user's password in the database.
	// Parameters:
	//   - userID: integer representing the user's unique ID
	//   - passwordHash: string containing the hashed version of the new password
	// Returns:
	//   - error: nil if update successful, otherwise contains the error message
	UpdatePassword(userID int, passwordHash string) error
}
