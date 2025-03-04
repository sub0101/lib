package errorss

import "fmt"

// Define error categories
type AppError struct {
	Code    int    // HTTP Status Code
	Message string // User-friendly error message
	Details string // Optional debugging details
}

func (e *AppError) Error() string {
	return fmt.Sprintf("[%d] %s: %s", e.Code, e.Message, e.Details)
}

func NotFound(message, details string) *AppError {
	return &AppError{Code: 404, Message: message, Details: details}
}

func BadRequest(message, details string) *AppError {
	return &AppError{Code: 400, Message: message, Details: details}
}

func InternalServerError(message, details string) *AppError {
	return &AppError{Code: 500, Message: message, Details: details}
}

func UnAuthrized(message, details string) *AppError {
	return &AppError{Code: 403, Message: message, Details: details}
}
func Forbidden(message, details string) *AppError {
	return &AppError{Code: 403, Message: message, Details: details}
}
