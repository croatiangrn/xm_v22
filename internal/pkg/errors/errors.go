package errors

import "fmt"

const (
	NotFoundMessage       = "%s with ID %s not found"
	BadRequestMessage     = "validation error on field %s: %s"
	InternalServerMessage = "internal server error during %s: %v"
)

type NotFoundError struct {
	Entity string
	ID     string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf(NotFoundMessage, e.Entity, e.ID)
}

func NewNotFoundError(entity, id string) *NotFoundError {
	return &NotFoundError{
		Entity: entity,
		ID:     id,
	}
}

type BadRequestError struct {
	Field   string
	Message string
}

func (e *BadRequestError) Error() string {
	return fmt.Sprintf(BadRequestMessage, e.Field, e.Message)
}

func NewBadRequestError(field, message string) *BadRequestError {
	return &BadRequestError{
		Field:   field,
		Message: message,
	}
}

type InternalServerError struct {
	Operation string
	Err       error
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf(InternalServerMessage, e.Operation, e.Err)
}

func (e *InternalServerError) Unwrap() error {
	return e.Err
}

func NewInternalServerError(operation string, err error) *InternalServerError {
	return &InternalServerError{
		Operation: operation,
		Err:       err,
	}
}
