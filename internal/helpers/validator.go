package helpers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func DecodeAndValidate(r *http.Request, dst any) ([]ValidationError, error) {
	// Decode JSON body
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		return nil, errors.New("invalid JSON body")
	}

	// Validate struct tags
	err := validate.Struct(dst)
	if err == nil {
		return nil, nil
	}

	var ValidationErrors []ValidationError
	for _, e := range err.(validator.ValidationErrors) {
		ValidationErrors = append(ValidationErrors, ValidationError{
			Field:   e.Field(),
			Message: validationMessage(e),
		})
	}

	return ValidationErrors, nil
}

func validationMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return e.Field() + " is required"
	case "min":
		return e.Field() + " must be at least " + e.Param() + " characters"
	case "max":
		return e.Field() + " must be at least " + e.Param() + " characters"
	default:
		return e.Field() + " is invalid"
	}
}
