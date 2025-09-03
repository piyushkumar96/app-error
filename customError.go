package errors

// CustomErr represents a structured custom error
type CustomErr struct {
	Code      string // The most important error code for API response
	Message   string // Human-readable error message
	Retryable bool   // Indicates whether the error is retryable
}

// GetCustomErr creates a new instance of CustomErr
func GetCustomErr(code, msg string, retryable bool) *CustomErr {
	return &CustomErr{
		Code:      code,
		Message:   msg,
		Retryable: retryable,
	}
}
