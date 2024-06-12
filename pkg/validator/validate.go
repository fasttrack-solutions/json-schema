package validator

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

func validatePayload(key string, payload []byte, schemas map[string]gojsonschema.JSONLoader) ([]Error, error) {
	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	validatorSchema, exists := schemas[key]
	if !exists {
		return nil, fmt.Errorf("no validator schema found for payload type %s", key)
	}

	errors, err := validateSchema(payload, validatorSchema)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %v", key, err)
	}

	return errors, nil
}

func validateSchema(payload []byte, validatorSchema gojsonschema.JSONLoader) ([]Error, error) {
	documentLoader := gojsonschema.NewBytesLoader(payload)

	validationResult, err := gojsonschema.Validate(validatorSchema, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("validating paylod: %v", err)
	}

	if validationResult.Valid() {
		return nil, nil
	}

	returnErrors := make([]Error, len(validationResult.Errors()))
	for i, error := range validationResult.Errors() {
		returnErrors[i] = Error{
			Path:  error.Context().String(),
			Error: error.Description(),
		}
	}

	return returnErrors, nil
}

func isValidType(t string, array []string) bool {
	return slices.Contains(array, strings.ToLower(t))
}
