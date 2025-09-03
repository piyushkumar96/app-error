package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	ae "github.com/piyushkumar96/app-error"
)

// Common error codes using CustomErr
var (
	OnDBPingFailure = ae.GetCustomErr(
		"ERR_SVC_1001",
		"database is not reachable",
		false)
	OnHTTPRequestMethodNotAllowed = ae.GetCustomErr(
		"ERR_SVC_1002",
		"method not allowed",
		false)
	UnexpectedError = ae.GetCustomErr(
		"ERR_SVC_1003",
		"unexpected error occurred",
		false)
	UnAuthorizedStaticToken = ae.GetCustomErr(
		"ERR_SVC_1004",
		"access token is missing or invalid",
		false)
	// Retryable error example
	ServiceTemporarilyUnavailable = ae.GetCustomErr(
		"ERR_SVC_1005",
		"service temporarily unavailable",
		true)
)

// Example 1: Database connection error
func ExampleDatabaseError() {
	ctx := context.Background()

	// Simulate a database ping failure
	dbErr := sql.ErrConnDone

	// Create an AppError with custom error details
	appErr := ae.GetAppErr(ctx, dbErr, OnDBPingFailure, http.StatusServiceUnavailable)

	fmt.Printf("Error Code: %s\n", appErr.GetErrCode())
	fmt.Printf("Error Message: %s\n", appErr.GetMsg())
	fmt.Printf("HTTP Code: %d\n", appErr.GetHTTPCode())
	fmt.Printf("Underlying Error: %s\n", appErr.GetErr())
}

// Example 2: HTTP method not allowed error
func ExampleHTTPMethodError() {
	ctx := context.Background()

	// Simulate an HTTP method error
	httpErr := fmt.Errorf("POST method not allowed on this endpoint")

	// Create AppError with method not allowed custom error
	appErr := ae.GetAppErr(ctx, httpErr, OnHTTPRequestMethodNotAllowed, http.StatusMethodNotAllowed)

	// Add additional metadata
	appErr.SetData(map[string]interface{}{
		"allowed_methods": []string{"GET", "PUT"},
		"endpoint":        "/api/users",
	})

	fmt.Printf("Error Code: %s\n", appErr.GetErrCode())
	fmt.Printf("Error Message: %s\n", appErr.GetMsg())
	fmt.Printf("HTTP Code: %d\n", appErr.GetHTTPCode())
	fmt.Printf("Additional Data: %v\n", appErr.GetData())
}

// Example 3: Authorization error
func ExampleAuthorizationError() {
	ctx := context.Background()

	// Simulate an unauthorized access attempt
	authErr := fmt.Errorf("invalid token provided")

	// Create AppError with unauthorized custom error
	appErr := ae.GetAppErr(ctx, authErr, UnAuthorizedStaticToken, http.StatusUnauthorized)

	// Chain additional error codes if needed
	appErr.AddErrCode("ERR_AUTH_TOKEN_EXPIRED")

	fmt.Printf("Primary Error Code: %s\n", appErr.GetErrCode())
	fmt.Printf("All Error Codes: %v\n", appErr.GetErrCodes())
	fmt.Printf("Error Message: %s\n", appErr.GetMsg())
}

// Example 4: Chaining multiple errors
func ExampleErrorChaining() {
	ctx := context.Background()

	// Simulate multiple error scenarios
	originalErr := fmt.Errorf("connection timeout")

	// Start with unexpected error
	appErr := ae.GetAppErr(ctx, originalErr, UnexpectedError, http.StatusInternalServerError)

	// Add more specific error codes as investigation continues
	appErr.AddErrCode("ERR_NETWORK_TIMEOUT").
		AddErrCode("ERR_DOWNSTREAM_SERVICE_UNAVAILABLE")

	// Update message with more specific information
	appErr.SetMsg("downstream service connection timeout")

	// Add debugging metadata
	appErr.SetData(map[string]interface{}{
		"service":        "payment-service",
		"timeout_ms":     5000,
		"retry_attempts": 3,
	})

	fmt.Printf("Error Evolution:\n")
	fmt.Printf("All Error Codes: %v\n", appErr.GetErrCodes())
	fmt.Printf("Final Message: %s\n", appErr.GetMsg())
	fmt.Printf("Debug Data: %v\n", appErr.GetData())
}

// Example 5: Using retryable errors
func ExampleRetryableError() {
	ctx := context.Background()

	// Simulate a temporary service unavailability
	serviceErr := fmt.Errorf("service overloaded")

	// Create AppError with retryable custom error
	appErr := ae.GetAppErr(ctx, serviceErr, ServiceTemporarilyUnavailable, http.StatusServiceUnavailable)

	// Check if error is retryable
	isRetryable := appErr.CustomErr.Retryable

	fmt.Printf("Error Code: %s\n", appErr.GetErrCode())
	fmt.Printf("Error Message: %s\n", appErr.GetMsg())
	fmt.Printf("Is Retryable: %t\n", isRetryable)

	if isRetryable {
		fmt.Println("Client can retry this request after some delay")
	}
}

// Example 6: Creating custom errors without predefined constants
func ExampleAdHocCustomError() {
	ctx := context.Background()

	// Create a custom error on the fly
	customErr := ae.GetCustomErr(
		"ERR_VALIDATION_001",
		"invalid email format provided",
		false)

	validationErr := fmt.Errorf("email validation failed for: user@invalid")

	appErr := ae.GetAppErr(ctx, validationErr, customErr, http.StatusBadRequest)

	// Add validation-specific metadata
	appErr.SetData(map[string]interface{}{
		"field":           "email",
		"provided_value":  "user@invalid",
		"expected_format": "user@domain.com",
	})

	fmt.Printf("Custom Error Code: %s\n", appErr.GetErrCode())
	fmt.Printf("Custom Message: %s\n", appErr.GetMsg())
	fmt.Printf("Validation Details: %v\n", appErr.GetData())
}

func main() {
	fmt.Println("=== App Error Examples ===")

	fmt.Println("1. Database Error Example:")
	ExampleDatabaseError()
	fmt.Println()

	fmt.Println("2. HTTP Method Error Example:")
	ExampleHTTPMethodError()
	fmt.Println()

	fmt.Println("3. Authorization Error Example:")
	ExampleAuthorizationError()
	fmt.Println()

	fmt.Println("4. Error Chaining Example:")
	ExampleErrorChaining()
	fmt.Println()

	fmt.Println("5. Retryable Error Example:")
	ExampleRetryableError()
	fmt.Println()

	fmt.Println("6. Ad-hoc Custom Error Example:")
	ExampleAdHocCustomError()
}
