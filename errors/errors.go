package errors

import "fmt"

// BasicError is general purpose error
type BasicError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (e BasicError) Error() string {
	return e.Message
}

// NewBasicError returns new BasicError instance
func NewBasicError(code, message string) BasicError {
	return BasicError{
		code, message,
	}
}

// RealmError is generic realm error
type RealmError struct {
	BasicError
}

func (e RealmError) Error() string {
	return e.Message
}

// NewRealmError returns new RealmError instance
func NewRealmError(message string) RealmError {
	return RealmError{
		BasicError: BasicError{
			"RealmError", message,
		},
	}
}

// ServiceError is generic service error
type ServiceError struct {
	BasicError
}

func (e ServiceError) Error() string {
	return e.Message
}

// NewServiceError returns new ServiceError instance
func NewServiceError(message string) ServiceError {
	return ServiceError{
		BasicError: BasicError{
			"ServiceError", message,
		},
	}
}

// AuthError is generic auth error
type AuthError struct {
	BasicError
}

func (e AuthError) Error() string {
	return e.Message
}

// NewAuthError returns new ServiceError instance
func NewAuthError(message string) AuthError {
	return AuthError{
		BasicError: BasicError{
			"AuthError", message,
		},
	}
}

// ValidationError is error used for all validation related errors
type ValidationError struct {
	*BasicError
	Fields []*ValidationErrorField `json:"fields,omitempty"`
}

// ValidationErrorField is a type to contain information on field error
type ValidationErrorField struct {
	Name    string `json:"name"`    // name of the field
	Message string `json:"message"` // error message related to the field
}

// NewValidationError returns new ValidationError
func NewValidationError() *ValidationError {
	return &ValidationError{
		BasicError: &BasicError{
			Code:    "ValidationError",
			Message: "Validation error",
		},
		Fields: []*ValidationErrorField{},
	}
}

// Error implementations of errors.Error interface
func (e ValidationError) Error() string {
	return e.Message
}

// ClearFieldErrors clears all field errors
func (e *ValidationError) ClearFieldErrors() {
	e.Fields = e.Fields[:0]
}

// FieldError sets field error
func (e *ValidationError) FieldError(name, message string) {
	for _, f := range e.Fields {
		if f.Name == name {
			f.Message = message
			return
		}
	}
	e.Fields = append(e.Fields, &ValidationErrorField{
		Name:    name,
		Message: message,
	})
}

// GetFieldError returns field error
func (e *ValidationError) GetFieldError(name string) *ValidationErrorField {
	for _, f := range e.Fields {
		if f.Name == name {
			return f
		}
	}
	return nil
}

// HasFieldErrors returns true when has field error
func (e *ValidationError) HasFieldErrors() bool {
	return len(e.Fields) > 0
}

// HasFieldError returns true when has field error
func (e *ValidationError) HasFieldError(field string) bool {
	return e.GetFieldError(field) != nil
}

// FieldInvalid set field with invalid error message
func (e *ValidationError) FieldInvalid(field string) {
	e.FieldError(field, fmt.Sprintf("field %s is invalid", field))
}

// FieldRequired set field with required error message
func (e *ValidationError) FieldRequired(field string) {
	e.FieldError(field, fmt.Sprintf("field %s is required", field))
}
