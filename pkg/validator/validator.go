package validator

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"slices"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/*
var schemasFileSystem embed.FS

const eventSchemasBasePath = "schemas/real-time-events"
const operatorAPISchemasBasePath = "schemas/operator-api/get"

var notificationTypes = []string{
	"bonus",
	"casino",
	"custom",
	"game",
	"login_v2",
	"lottery_v2",
	"payment",
	"sportsbook",
	"user_balances_update",
	"user_block_v2",
	"user_consents_v2",
	"user_create_v2",
	"user_update_v2",
}

var operatorAPIEndpoints = []string{
	"user_details",
	"user_blocks",
	"user_consents",
}

type Client struct {
	eventSchemas       map[string]gojsonschema.JSONLoader
	operatorAPISchemas map[string]gojsonschema.JSONLoader

	publicEventSchemas       map[string]map[string]interface{}
	publicOperatorAPISchemas map[string]map[string]interface{}
}

func New() (*Client, error) {
	publicEventSchemas := make(map[string]map[string]interface{}, len(notificationTypes))
	eventSchemas := make(map[string]gojsonschema.JSONLoader, len(notificationTypes))
	for _, nt := range notificationTypes {
		filename := fmt.Sprintf("%s/%s.schema.json", eventSchemasBasePath, nt)
		fileData, err := schemasFileSystem.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("reading file %s from embedded schemas file system: %v", nt, err)
		}

		eventSchemas[nt] = gojsonschema.NewBytesLoader(fileData)

		var rawSchema map[string]interface{}
		err = json.Unmarshal(fileData, &rawSchema)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		publicEventSchemas[nt] = rawSchema
	}

	publicOperatorAPISchemas := make(map[string]map[string]interface{}, len(operatorAPIEndpoints))
	operatorAPISchemas := make(map[string]gojsonschema.JSONLoader, len(operatorAPIEndpoints))
	for _, endpoint := range operatorAPIEndpoints {
		filename := fmt.Sprintf("%s/%s.schema.json", operatorAPISchemasBasePath, endpoint)
		fileData, err := schemasFileSystem.ReadFile(filename)
		if err != nil {
			return nil, fmt.Errorf("reading file %s from embedded schemas file system: %v", endpoint, err)
		}

		operatorAPISchemas[endpoint] = gojsonschema.NewBytesLoader(fileData)

		var rawSchema map[string]interface{}
		err = json.Unmarshal(fileData, &rawSchema)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		publicOperatorAPISchemas[endpoint] = rawSchema
	}

	return &Client{
		eventSchemas:       eventSchemas,
		publicEventSchemas: publicEventSchemas,

		operatorAPISchemas:       operatorAPISchemas,
		publicOperatorAPISchemas: publicOperatorAPISchemas,
	}, nil
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetEventSchemas() map[string]map[string]interface{} {
	return c.publicEventSchemas
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetOperatorAPISchemas() map[string]map[string]interface{} {
	return c.publicOperatorAPISchemas
}

func isValidNotificationType(t string) bool {
	return slices.Contains(notificationTypes, strings.ToLower(t))
}

func isValidEndpoint(t string) bool {
	return slices.Contains(operatorAPIEndpoints, strings.ToLower(t))
}

// ValidateEvent takes in notificationType which can be the notification type for the real time feed event, and payload which is the json payload.
// Allowed notification types are: bonus, casino, custom, game, login_v2, lottery_v2, payment, sportsbook, user_balances_update, user_block_v2, user_consents_v2, user_create_v2, user_update_v2
func (c *Client) ValidateEvent(notificationType string, payload []byte) ([]Error, error) {
	if notificationType == "" {
		return nil, errors.New("no payload type provided")
	}

	if !isValidNotificationType(notificationType) {
		return nil, fmt.Errorf("%s is an invalid payload type", notificationType)
	}

	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	validatorSchema, exists := c.eventSchemas[notificationType]
	if !exists {
		return nil, fmt.Errorf("no validator schema found for payload type %s", notificationType)
	}

	errors, err := validate(validatorSchema, payload)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %v", notificationType, err)
	}

	return errors, nil
}

func (c *Client) ValidateOperatorAPIResponse(endpoint string, payload []byte) ([]Error, error) {
	if endpoint == "" {
		return nil, errors.New("no endpoint provided")
	}

	if !isValidEndpoint(endpoint) {
		return nil, fmt.Errorf("%s is an invalid endpoint", endpoint)
	}

	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	validatorSchema, exists := c.operatorAPISchemas[endpoint]
	if !exists {
		return nil, fmt.Errorf("no validator schema found for endpoint %s", endpoint)
	}

	errors, err := validate(validatorSchema, payload)
	if err != nil {
		return nil, fmt.Errorf("validating operator api endpoint %s response json: %v", endpoint, err)
	}

	return errors, nil
}

func validate(validatorSchema gojsonschema.JSONLoader, payload []byte) ([]Error, error) {
	documentLoader := gojsonschema.NewBytesLoader(payload)

	validationResult, err := gojsonschema.Validate(validatorSchema, documentLoader)
	if err != nil {
		return nil, fmt.Errorf("validating paylod: %v", err)
	}

	if validationResult.Valid() {
		return nil, nil
	}

	returnErrors := make([]Error, len(validationResult.Errors()))
	for i, e := range validationResult.Errors() {
		returnErrors[i] = Error{
			Path:  e.Context().String(),
			Error: e.Description(),
		}
	}

	return returnErrors, nil
}

type Error struct {
	Path  string
	Error string
}
