package apperror

import (
	"fmt"
	"sort"
	"strings"
)

type AppError struct {
	Code       string
	Message    string
	Type       string
	Details    string
	HTTPStatus int
}

func (e AppError) Error() string {
	if e.Details != "" {
		return fmt.Sprintf("[%s] %s (%s): %s", e.Code, e.Message, e.Type, e.Details)
	}
	return fmt.Sprintf("[%s] %s (%s)", e.Code, e.Message, e.Type)
}

func New(code, errType, message string, httpStatus int, details ...string) AppError {
	err := AppError{
		Code:       code,
		Message:    message,
		Type:       errType,
		HTTPStatus: httpStatus,
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

type Errors []AppError

func (es Errors) Error() string {
	var errs []string
	for _, e := range es {
		errs = append(errs, e.Error())
	}
	sort.Strings(errs)
	return strings.Join(errs, ";")
}
