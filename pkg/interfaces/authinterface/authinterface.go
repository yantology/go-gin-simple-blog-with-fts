package authinterface

import (
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/authmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

// AuthInterface defines the methods for user authentication operations.
type AuthInterface interface {
	// CreateUser inserts a new user into the database.
	// Parameters:
	//   - username: unique identifier for the user
	//   - email: user's email address
	//   - password: user's password credential
	//
	// Returns:
	//   - *customerror.CustomError: nil if successful, error details if failed
	CreateUser(username string, email string, password string) *customerror.CustomError

	// GetUserByEmail retrieves user information using their email address.
	// Parameters:
	//   - email: string containing the user's email address
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByEmail(email string) (*authmodels.User, *customerror.CustomError)

	// GetUserByID retrieves user information using their unique identifier.
	// Parameters:
	//   - id: integer representing the user's unique ID in the database
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByID(id int) (*authmodels.User, *customerror.CustomError)

	// GetUserByUsername retrieves user information using their username.
	// Parameters:
	//   - username: string containing the user's chosen username
	// Returns:
	//   - *User: user object if found, nil if user doesn't exist
	//   - error: nil if operation successful, otherwise contains the error message
	GetUserByUsername(username string) (*authmodels.User, *customerror.CustomError)

	// CheckUsernameExists verifies if a username is already taken in the database.
	// Parameters:
	//   - username: string containing the username to check
	// Returns:
	//   - bool:
	//     * true: username is already registered by another user and cannot be used
	//     * false: username is available and can be used for registration
	//   - error: nil if check successful, otherwise contains database error details
	CheckUsernameExists(username string) (bool, *customerror.CustomError)

	// UpdatePassword changes a user's password in the database.
	// Parameters:
	//   - userID: integer representing the user's unique ID
	//   - passwordHash: string containing the hashed version of the new password
	// Returns:
	//   - error: nil if update successful, otherwise contains the error message
	UpdatePassword(userID int, passwordHash string) *customerror.CustomError
}
