package postgreserror

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/lib/pq"
	"github.com/yantology/go-gin-simple-blog-with-fts/pkg/utils/customerror"
)

// NewPostgresError creates a custom error from PostgreSQL errors
func NewPostgresError(err error) *customerror.CustomError {
	if err == nil {
		return nil
	}

	// Handle specific PostgreSQL errors
	if pqErr, ok := err.(*pq.Error); ok {
		switch pqErr.Code {
		case "23505": // unique_violation
			return customerror.NewCustomError(err, "Record already exists", http.StatusBadRequest)
		case "23503": // foreign_key_violation
			return customerror.NewCustomError(err, "Foreign key violation", http.StatusBadRequest)
		case "22001": // string_data_right_truncation
			return customerror.NewCustomError(err, "String data is too long", http.StatusBadRequest)
		}
	}

	// Generic database error
	if errors.Is(err, sql.ErrNoRows) {
		return customerror.NewCustomError(err, "Record not found", http.StatusNotFound)
	}

	// Default error handling
	return customerror.NewCustomError(err, "Database error", http.StatusInternalServerError)
}
