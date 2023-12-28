package domain

// NotFoundError is an error for not found.
type NotFoundError struct{}

func NewNotFoundError() *NotFoundError {
	return &NotFoundError{}
}

func (e *NotFoundError) Error() string {
	return ""
}

// AlreadyExistsError is an error for data that already exists.
type AlreadyExistsError struct {
	Message string
}

func NewAlreadyExistsError(message string) *AlreadyExistsError {
	return &AlreadyExistsError{Message: message}
}

func (e *AlreadyExistsError) Error() string {
	return e.Message
}

// ValidationError is an error for invalid request.
type ValidationError struct{}

func NewValidationError() *ValidationError {
	return &ValidationError{}
}

func (e *ValidationError) Error() string {
	return ""
}

type UnauthorizedError struct{}

func NewUnauthorizedError() *UnauthorizedError {
	return &UnauthorizedError{}
}

func (e *UnauthorizedError) Error() string {
	return ""
}

type ForbiddenError struct{}

func NewForbiddenError() *ForbiddenError {
	return &ForbiddenError{}
}

func (e *ForbiddenError) Error() string {
	return ""
}
