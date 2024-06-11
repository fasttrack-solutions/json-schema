package validator

import (
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// ValidateEvent takes in notificationType which can be the notification type for the real time feed event, and payload which is the json payload.
// Allowed notification types are: bonus, casino, custom, game, login_v2, lottery_v2, payment, sportsbook, user_balances_update, user_block_v2, user_consents_v2, user_create_v2, user_update_v2
func (c *Client) ValidateEvent(notificationType string, payload []byte) ([]Error, error) {
	if notificationType == "" {
		return nil, errors.New("no payload type provided")
	}

	if !isValidType(notificationType, notificationTypes) {
		return nil, fmt.Errorf("%s is an invalid notification type", notificationType)
	}

	errors, err := validatePayload(notificationType, payload, c.eventSchemas)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %v", notificationType, err)
	}

	return errors, nil
}

func (c *Client) ValidateOperatorAPIResponse(endpoint string, payload []byte) ([]Error, error) {
	if endpoint == "" {
		return nil, errors.New("no endpoint provided")
	}

	if !isValidType(endpoint, operatorAPIEndpoints) {
		return nil, fmt.Errorf("%s is an invalid endpoint", endpoint)
	}

	errors, err := validatePayload(endpoint, payload, c.operatorAPISchemas)
	if err != nil {
		return nil, fmt.Errorf("validating %s endpoint json payload: %v", endpoint, err)
	}

	return errors, nil
}

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
