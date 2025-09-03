package errors

import (
	"context"
)

// AppError represents a structured error with additional metadata
type AppError struct {
	ActualErr  error       // The actual underlying error
	CustomErr  *CustomErr  // Custom error details (code, message, etc.)
	ErrorCodes []string    // All error codes encountered during execution
	httpCode   int         // Corresponding HTTP error code
	data       interface{} // Additional data to include in the error response
}

// Error implements the error interface, returning the error message
func (e *AppError) Error() string {
	if e.ActualErr == nil {
		return ""
	}
	return e.ActualErr.Error()
}

// GetErr retrieves the underlying error
func (e *AppError) GetErr() error {
	return e.ActualErr
}

// SetErr sets the underlying error and returns it
func (e *AppError) SetErr(err error) error {
	e.ActualErr = err
	return e.ActualErr
}

// GetMsg retrieves the custom error message
func (e *AppError) GetMsg() string {
	return e.CustomErr.Message
}

// SetMsg updates the custom error message and returns the AppError
func (e *AppError) SetMsg(msg string) *AppError {
	e.CustomErr.Message = msg
	return e
}

// GetHTTPCode retrieves the HTTP status code
func (e *AppError) GetHTTPCode() int {
	return e.httpCode
}

// SetHTTPCode updates the HTTP status code and returns the AppError
func (e *AppError) SetHTTPCode(httpCode int) *AppError {
	e.httpCode = httpCode
	return e
}

// GetErrCode retrieves the primary error code
func (e *AppError) GetErrCode() string {
	return e.CustomErr.Code
}

// SetErrCode updates the primary error code and returns the AppError
func (e *AppError) SetErrCode(code string) *AppError {
	e.CustomErr.Code = code
	return e
}

// GetErrCodes retrieves the list of all error codes
func (e *AppError) GetErrCodes() []string {
	return e.ErrorCodes
}

// AddErrCode appends an error code to the list and updates the primary code
func (e *AppError) AddErrCode(errorCode string) *AppError {
	if errorCode != "" {
		e.CustomErr.Code = errorCode
		e.ErrorCodes = append(e.ErrorCodes, errorCode)
	}
	return e
}

// GetData retrieves the additional metadata associated with the error
func (e *AppError) GetData() interface{} {
	return e.data
}

// SetData updates the metadata and returns the AppError
func (e *AppError) SetData(data interface{}) *AppError {
	e.data = data
	return e
}

// GetAppErr creates a new instance of AppError
func GetAppErr(ctx context.Context, err error, customErr *CustomErr, httpCode int, meta ...interface{}) *AppError {
	// Log the error trace for debugging
	AddTraceLog(ctx, err.Error())

	// Initialize the AppError structure
	appErr := &AppError{
		ActualErr:  err,
		CustomErr:  &CustomErr{},
		httpCode:   httpCode,
		ErrorCodes: []string{},
	}

	// Assign metadata if provided
	if len(meta) > 0 {
		appErr.data = meta[0]
	}

	// Populate custom error details if provided
	if customErr != nil {
		appErr.CustomErr.Code = customErr.Code
		appErr.CustomErr.Message = customErr.Message
		appErr.ErrorCodes = append(appErr.ErrorCodes, customErr.Code)
	}

	return appErr
}
