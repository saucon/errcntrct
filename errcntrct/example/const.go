// This file is generated using errcntrct tool.
// Check out for more info "https://github.com/Saucon/errcntrct"
package contract

import "errors"

const (
	ErrInvalidRequestFamily_const = "1000"
	ErrInvalidDateFormat_const = "1002"
	ErrEmailNotFound_const = "1003"
	ErrUnexpectedError_const = "9999"
)
var (
	ErrInvalidRequestFamily = errors.New(ErrInvalidRequestFamily_const)
	ErrInvalidDateFormat = errors.New(ErrInvalidDateFormat_const)
	ErrEmailNotFound = errors.New(ErrEmailNotFound_const)
	ErrUnexpectedError = errors.New(ErrUnexpectedError_const)
)
