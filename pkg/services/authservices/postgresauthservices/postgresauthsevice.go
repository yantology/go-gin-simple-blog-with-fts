package postgresauthservices

import (
	"database/sql"
	"time"

	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/models/authmodels"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror/postgreserror"
)

type PostgresAuthService struct {
	db *sql.DB
}

// NewPostgresAuthService creates a new instance of PostgresAuthService
func NewPostgresAuthService(db *sql.DB) *PostgresAuthService {

	return &PostgresAuthService{db: db}
}

// CreateUser creates a new user in the database.
// Parameters:
//   - r: CreateUserRequest containing username, email, and password
//
// Returns:
//   - error: nil if user created successfully, otherwise contains error details
func (s *PostgresAuthService) CreateUser(username string, email string, password string) *customerror.CustomError {

	query := `
        INSERT INTO users (username, email, password, created_at, updated_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id, created_at, updated_at`
	_, err := s.db.Exec(query, username, email, password, time.Now(), time.Now())
	if err != nil {
		return postgreserror.NewPostgresError(err)
	}
	return nil
}

// GetUserByEmail retrieves user information using their email address.
// Parameters:
//   - email: string containing the user's email address to search
//
// Returns:
//   - *User: contains user details if found, nil if no user exists with this email
//   - error: nil if query successful, otherwise contains database error details
func (s *PostgresAuthService) GetUserByEmail(email string) (*authmodels.User, *customerror.CustomError) {
	user := &authmodels.User{}
	query := `
        SELECT id, username, email, password, created_at, updated_at
        FROM users WHERE email = $1`

	err := s.db.QueryRow(query, email).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password,
			&user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	return user, nil
}

// GetUserByID retrieves user information using their unique identifier.
// Parameters:
//   - id: integer containing the user's unique database ID
//
// Returns:
//   - *User: contains user details if found, nil if no user exists with this ID
//   - error: nil if query successful, otherwise contains database error details
func (s *PostgresAuthService) GetUserByID(id int) (*authmodels.User, *customerror.CustomError) {
	user := &authmodels.User{}
	query := `
        SELECT id, username, email, password, created_at, updated_at
        FROM users WHERE id = $1`

	err := s.db.QueryRow(query, id).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password,
			&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	return user, nil
}

// GetUserByUsername retrieves user information using their username.
// Parameters:
//   - username: string containing the user's username to search
//
// Returns:
//   - *User: contains user details if found, nil if no user exists with this username
//   - error: nil if query successful, otherwise contains database error details
func (s *PostgresAuthService) GetUserByUsername(username string) (*authmodels.User, *customerror.CustomError) {
	user := &authmodels.User{}
	query := `
        SELECT id, username, email, password, created_at, updated_at
        FROM users WHERE username = $1`

	err := s.db.QueryRow(query, username).
		Scan(&user.ID, &user.Username, &user.Email, &user.Password,
			&user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		return nil, postgreserror.NewPostgresError(err)
	}
	return user, nil
}

// CheckUsernameExists verifies if a username is already registered in the database.
// Parameters:
//   - username: string containing the username to check
//
// Returns:
//   - bool:
//   - true: username is already taken by another user (cannot be used)
//   - false: username is available for registration (can be used)
//   - error: nil if check successful, otherwise contains database error details
func (s *PostgresAuthService) CheckUsernameExists(username string) (bool, *customerror.CustomError) {
	query := `SELECT EXISTS(SELECT 1 FROM users WHERE username = $1)`
	var exists bool
	err := s.db.QueryRow(query, username).Scan(&exists)
	if err != nil {
		return false, postgreserror.NewPostgresError(err)

	}
	return exists, nil
}

// UpdatePassword updates the password hash for a specific user.
// Parameters:
//   - userID: integer containing the user's unique database ID
//   - password: string containing the new password hash to store
//
// Returns:
//   - error: nil if password updated successfully, otherwise contains error details
func (s *PostgresAuthService) UpdatePassword(userID int, password string) *customerror.CustomError {
	query := `
        UPDATE users 
        SET password = $1, updated_at = CURRENT_TIMESTAMP
        WHERE id = $2`

	_, err := s.db.Exec(query, password, userID)
	if err != nil {
		return postgreserror.NewPostgresError(err)
	}
	return nil
}
