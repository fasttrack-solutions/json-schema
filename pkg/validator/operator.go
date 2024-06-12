package validator

import (
	"errors"
	"fmt"
)

const pathSchemasOperatorAPI = "schemas/operator-api/get"

var operatorAPIEndpoints = []string{
	"user_details",
	"user_blocks",
	"user_consents",
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
