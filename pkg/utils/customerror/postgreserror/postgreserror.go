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
			return &customerror.CustomError{
				HTTPCode: http.StatusConflict,
				Message:  "Record already exists",
				Original: err,
			}
		case "23503": // foreign_key_violation
			return &customerror.CustomError{
				HTTPCode: http.StatusBadRequest,
				Message:  "Related record not found",
				Original: err,
			}
		case "22001": // string_data_right_truncation
			return &customerror.CustomError{
				HTTPCode: http.StatusBadRequest,
				Message:  "Data too long for field",
				Original: err,
			}
		}
	}

	// Generic database error
	if errors.Is(err, sql.ErrNoRows) {
		return &customerror.CustomError{
			HTTPCode: http.StatusNotFound,
			Message:  "No records found",
			Original: err,
		}
	}

	// Default error handling
	return &customerror.CustomError{
		HTTPCode: http.StatusInternalServerError,
		Message:  "Database operation failed",
		Original: err,
	}
}
