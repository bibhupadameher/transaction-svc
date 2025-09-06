package apperror

import "fmt"

type AppError struct {
	Code    string
	Message string
	Type    string
	Details string
}

func (e *AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s (%s): %s", e.Code, e.Message, e.Type, e.Details)
	}
	return fmt.Sprintf("[%s] %s (%s)", e.Code, e.Message, e.Type)
}

func New(code, message, errType string, details ...string) *AppError {
	err := &AppError{
		Code:    code,
		Message: message,
		Type:    errType,
	}
	if len(details) > 0 {
		err.Details = details[0]
	}
	return err
}

func Is(err error, code string) bool {
	if appErr, ok := err.(*AppError); ok {
		return appErr.Code == code
	}
	return false
}
