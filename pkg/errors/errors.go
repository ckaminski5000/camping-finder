package errors

import "fmt"

type CustomError struct {
	StatusCode int
	Message    string
}

func CreateError(statusCode int, message string) *CustomError {
	return &CustomError{statusCode, message}
}

func (c CustomError) Error() string {
	return fmt.Sprintf("Status Code: %d, Error Message: %s", c.StatusCode, c.Message)
}
