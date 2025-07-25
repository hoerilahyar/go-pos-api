package errors

import "log"

type AppError struct {
	Code    string
	Message string
	Errors  error
}

// Error implements the error interface
func (e AppError) Error() string {
	return e.Message
}

// New creates a new AppError
func New(code, message string) AppError {
	return AppError{Code: code, Message: message}
}

func Get(appErr AppError, err error) error {
	if err != nil {
		log.Printf("AppError [%s]: %s | Cause: %v", appErr.Code, appErr.Message, err)
	} else {
		log.Printf("AppError [%s]: %s", appErr.Code, appErr.Message)
	}

	appErr.Errors = err

	return appErr
}
