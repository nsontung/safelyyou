package syerror

import "errors"

// define custom error types for better error handling in the API

var (
	ErrDeviceNotFound = errors.New("Device not found")
	ErrServerError    = errors.New("Server error")
)
