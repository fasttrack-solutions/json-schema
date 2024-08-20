package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const (
	pathSchemasOperatorAPI = "schemas/operator-api"
	pathSchemasRealTime    = "schemas/real-time-events"
)

// ValidateRealTimeEvent takes in notificationType which can be the notification type for the real time feed event, and payload which is the json payload.
// Allowed notification types are: bonus, casino, custom, game, login_v2, lottery_v2, payment, sportsbook, user_balances_update, user_block_v2, user_consents_v2, user_create_v2, user_update_v2
func (c *Client) ValidateRealTimeEvent(notificationType string, payload []byte) ([]ValidationError, error) {
	return c.validatePayload(notificationType, payload, c.realTimeSchemas)
}

// ValidateOperatorAPIResponse validates the given payload against the operator API schemas.
// Allowed operator API endpoints are: user_details, user_blocks, user_consents
func (c *Client) ValidateOperatorAPIResponse(endpoint string, payload []byte) ([]ValidationError, error) {
	return c.validatePayload(endpoint, payload, c.operatorAPISchemas)
}

func (c *Client) validatePayload(key string, payload []byte, schemasMap map[string]gojsonschema.Schema) ([]ValidationError, error) {
	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	schemaFile, exists := schemasMap[strings.ToLower(key)]
	if !exists {
		return nil, fmt.Errorf("invalid payload type %s", key)
	}

	errors, err := performValidation(payload, schemaFile)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %v", key, err)
	}

	return errors, nil
}

func performValidation(payload []byte, validationSchema gojsonschema.Schema) ([]ValidationError, error) {
	schemaBytesLoader := gojsonschema.NewBytesLoader(payload)

	validationResult, err := validationSchema.Validate(schemaBytesLoader)
	if err != nil {
		return nil, fmt.Errorf("validating paylod: %v", err)
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
