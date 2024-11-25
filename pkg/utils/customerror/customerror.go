package customerror

type CustomError struct {
	HTTPCode int
	Message  string
	Original error
}

// Error implements the error interface
func (ce *CustomError) Error() string {
	return ce.Message
}

// Helper method to extract original database error message
func (ce *CustomError) OriginalMessage() string {
	if ce.Original != nil {
		return ce.Original.Error()
	}
	return ""
}

// Helper method to extract original database error code
func (ce *CustomError) OriginalCode() int {
	if ce.Original != nil {
		return ce.HTTPCode
	}
	return 0
}
