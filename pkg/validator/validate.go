package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const (
	pathSchemasOperatorAPI = "schemas/operator-api/get"
	pathSchemasRealTime    = "schemas/real-time-events"
)

// ErrSchemaNotFound is returned when no JSON schema is registered for the
// given payload type or endpoint. Callers can detect it with errors.Is to
// fall back gracefully instead of treating the payload as unprocessable.
var ErrSchemaNotFound = errors.New("schema not found")

// ValidateRealTimeEvent validates payload against the schema for the given notificationType.
// See getNotificationTypes() in validator.go for the full list of supported types.
func (c *Client) ValidateRealTimeEvent(notificationType string, payload []byte) ([]ValidationError, error) {
	return validatePayload(notificationType, payload, c.realTimeSchemas)
}

// ValidateOperatorAPIResponse validates the given payload against the operator API schemas.
// Allowed operator API endpoints are: user_details, user_blocks, user_consents
func (c *Client) ValidateOperatorAPIResponse(endpoint string, payload []byte) ([]ValidationError, error) {
	return validatePayload(endpoint, payload, c.operatorAPISchemas)
}

func validatePayload(key string, payload []byte, schemasMap map[string]gojsonschema.JSONLoader) ([]ValidationError, error) {
	key = strings.ToLower(key)
	validationSchema := schemasMap[key]
	if key == "" || validationSchema == nil {
		return nil, fmt.Errorf("invalid payload type %s: %w", key, ErrSchemaNotFound)
	}

	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	errors, err := performValidation(payload, validationSchema)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %w", key, err)
	}

	return errors, nil
}

func performValidation(payload []byte, validationSchema gojsonschema.JSONLoader) ([]ValidationError, error) {
	schemaBytesLoader := gojsonschema.NewBytesLoader(payload)

	validationResult, err := gojsonschema.Validate(validationSchema, schemaBytesLoader)
	if err != nil {
		return nil, fmt.Errorf("validating payload: %w", err)
	}

	// If the payload is valid, return no errors
	if validationResult.Valid() {
		return nil, nil
	}

	// If the payload is invalid, return the validation errors
	returnErrors := make([]ValidationError, len(validationResult.Errors()))
	for i, error := range validationResult.Errors() {
		returnErrors[i] = ValidationError{
			Path:  error.Context().String(),
			Error: error.Description(),
		}
	}

	return returnErrors, nil
}
