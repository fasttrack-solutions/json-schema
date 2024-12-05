package validator

import (
	"errors"
)

type ValidationError struct {
	Path  string
	Error string
}

func (ve *ValidationError) Validate() error {
	if ve.Path == "" {
		return errors.New("validation error path is empty")
	}

	if ve.Error == "" {
		return errors.New("validation error message is empty")
	}

	return nil
}
