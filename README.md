# App Error - Structured Error Handling for Go Applications

A comprehensive and flexible error handling package designed for Go applications that need structured, traceable, and contextual error management. This package provides a robust foundation for building reliable services with proper error propagation, logging, and debugging capabilities.

## Overview

App Error is built around the principle that effective error handling requires more than just error messages. Modern applications need structured errors that carry context, support tracing, provide HTTP status mappings, and enable proper debugging workflows. This package delivers all of these capabilities while maintaining simplicity and performance.

## Key Features

### Structured Error Management
- **Hierarchical Error Structure**: Combines actual errors with custom error definitions and metadata
- **Error Code Management**: Support for primary error codes and error code chains
- **HTTP Status Integration**: Automatic mapping between errors and appropriate HTTP status codes
- **Flexible Metadata**: Attach arbitrary data to errors for debugging and logging purposes

### Context-Aware Error Tracing
- **Trace Logging**: Automatic trace collection as errors propagate through your application
- **Context Integration**: Seamless integration with Go's context package for request tracing
- **Error Evolution Tracking**: Monitor how errors change as they move through different layers

### Custom Error Definitions
- **Predefined Error Types**: Create reusable error definitions with codes, messages, and retry policies
- **Dynamic Error Creation**: Build errors on-the-fly when predefined types aren't sufficient
- **Retry Policy Support**: Mark errors as retryable or non-retryable for client guidance

### Developer-Friendly API
- **Method Chaining**: Fluent interface for building complex error structures
- **Immutable Operations**: Safe error manipulation without side effects
- **Type Safety**: Strong typing throughout the API to prevent common mistakes

## Core Components

### AppError Structure
The central error type that encapsulates all error information including the underlying error, custom error details, error code chains, HTTP status codes, and additional metadata.

### CustomErr Structure
Represents predefined error templates with error codes, human-readable messages, and retry policies. These serve as blueprints for creating consistent error responses across your application.

### TraceMeta Structure
Manages error tracing information including trace logs, error evolution history, and identifier mappings for debugging and monitoring purposes.

## Installation

Add the package to your Go module:

```
go get github.com/piyushkumar96/app-error
```

Import the package in your Go files:

```
import ae "github.com/piyushkumar96/app-error"
```

## API Documentation

### Creating Custom Error Definitions

**GetCustomErr Function**
Creates a new custom error definition with the specified error code, message, and retry policy. This function is typically used to define error constants that can be reused throughout your application.

Parameters:
- `code`: Unique error identifier (string)
- `message`: Human-readable error description (string)
- `retryable`: Whether the error condition can be retried (boolean)

Returns a CustomErr pointer that can be used with GetAppErr.

### Creating Application Errors

**GetAppErr Function**
Constructs a complete application error by combining an actual error with custom error details, HTTP status code, and optional metadata.

Parameters:
- `ctx`: Request context for tracing (context.Context)
- `err`: The underlying error that occurred (error)
- `customErr`: Custom error definition (*CustomErr)
- `httpCode`: Appropriate HTTP status code (int)
- `meta`: Optional metadata (variadic interface{})

Returns an AppError pointer with all error information properly structured.

### AppError Methods

**Error Retrieval Methods**
- `Error()`: Returns the underlying error message (implements error interface)
- `GetErr()`: Retrieves the actual underlying error
- `GetMsg()`: Returns the custom error message
- `GetErrCode()`: Gets the primary error code
- `GetErrCodes()`: Returns all error codes in the chain
- `GetHTTPCode()`: Retrieves the HTTP status code
- `GetData()`: Accesses attached metadata

**Error Modification Methods**
- `SetErr(error)`: Updates the underlying error
- `SetMsg(string)`: Modifies the custom error message
- `SetErrCode(string)`: Changes the primary error code
- `SetHTTPCode(int)`: Updates the HTTP status code
- `SetData(interface{})`: Attaches or updates metadata
- `AddErrCode(string)`: Appends an error code to the chain

All modification methods return the AppError instance to enable method chaining.

### Context and Tracing

**AddTraceLog Function**
Adds error information to the trace log stored in the request context. This function automatically captures error details for debugging and monitoring purposes.

Parameters:
- `ctx`: Request context containing trace information (context.Context)
- `errorMsg`: Error message to add to the trace (string)

Returns the updated TraceMeta or nil if context is invalid.

## Usage Patterns

### Basic Error Creation
Start with a simple error and enhance it with custom error details and appropriate HTTP status codes. This pattern is suitable for straightforward error scenarios where you need structured error information.

### Error Evolution and Chaining
Begin with a general error and progressively add more specific error codes as your understanding of the problem improves. This pattern is excellent for debugging complex issues and providing detailed error context.

### Predefined Error Constants
Define common errors as package-level constants to ensure consistency across your application. This approach reduces duplication and provides a centralized error catalog.

### Metadata Attachment
Include additional debugging information such as request parameters, system state, or configuration details that can help with troubleshooting.

### Retryable Error Handling
Use the retry policy feature to guide client behavior, especially in distributed systems where temporary failures are common.

## Best Practices

### Error Code Conventions
- Use consistent prefixes for different service layers or modules
- Include severity levels in error codes when appropriate
- Maintain a centralized registry of error codes to prevent conflicts
- Use semantic versioning principles for error code evolution

### Message Guidelines
- Write clear, actionable error messages for end users
- Include technical details in metadata rather than user-facing messages
- Use consistent language and terminology across all error messages
- Avoid exposing internal implementation details in error messages

### HTTP Status Code Mapping
- Follow HTTP status code specifications strictly
- Use 4xx codes for client errors and 5xx codes for server errors
- Be consistent in status code usage across similar error types
- Consider the impact on API consumers when choosing status codes

### Context Management
- Always pass context through your error handling chain
- Use context cancellation to prevent resource leaks in error scenarios
- Store request-specific information in context for error tracing
- Avoid storing large objects in context to prevent memory issues

### Metadata Usage
- Include only essential debugging information in metadata
- Use structured data formats for complex metadata
- Consider the performance impact of large metadata objects
- Sanitize sensitive information before including in metadata

### Testing Considerations
- Test error scenarios as thoroughly as success cases
- Verify error code consistency across different code paths
- Validate HTTP status code mappings in integration tests
- Mock external dependencies to test error handling paths

## Performance Considerations

### Memory Management
The package is designed to minimize memory allocations during error creation and manipulation. Error objects are lightweight and metadata is stored efficiently.

### Concurrency Safety
All operations on error objects are safe for concurrent access. Context-based tracing handles concurrent requests appropriately without race conditions.

### Trace Overhead
Error tracing adds minimal overhead to request processing. Trace data is collected efficiently and stored in a compact format.

## Error Handling Strategies

### Layered Error Handling
Structure your application in layers with each layer adding appropriate context to errors as they propagate upward. Lower layers focus on technical details while upper layers add business context.

### Circuit Breaker Integration
Use the retry policy feature to integrate with circuit breakers and other reliability patterns. Mark errors as retryable or non-retryable based on the underlying failure type.

### Graceful Degradation
Design error responses that allow clients to degrade gracefully when certain features are unavailable. Use appropriate HTTP status codes and error messages to guide client behavior.

### Monitoring and Alerting
Leverage error codes and metadata for monitoring system health. Set up alerting based on error code patterns and frequencies.

## Contributing

We welcome contributions to improve this package. Please follow these guidelines:

### Development Setup
- Ensure Go 1.20 or later is installed
- Run tests before submitting changes
- Follow Go coding conventions and best practices
- Add appropriate documentation for new features

### Testing Requirements
- Write unit tests for all new functionality
- Ensure existing tests pass after changes
- Add integration tests for complex features
- Include edge case testing for error scenarios

### Documentation Updates
- Update this README for significant changes
- Add inline code documentation for new functions
- Include usage examples for new features
- Update API documentation as needed

## License

This project is licensed under the terms specified in the LICENSE file. Please review the license before using this package in your projects.

## Support and Feedback

For questions, bug reports, or feature requests, please use the appropriate channels provided by your organization. We value your feedback and contributions to making this package better.

## Versioning

This package follows semantic versioning principles. Check the version tags for compatibility information and upgrade guidance.

## Changelog

Significant changes and new features are documented in version releases. Review the changelog when upgrading to understand breaking changes and new capabilities.
