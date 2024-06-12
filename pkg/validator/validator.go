package validator

import (
	"github.com/xeipuuv/gojsonschema"
)

type Error struct {
	Path  string
	Error string
}

type Client struct {
	eventSchemas        map[string]gojsonschema.JSONLoader
	eventSchemaRegistry map[string]map[string]interface{}

	operatorAPISchemas         map[string]gojsonschema.JSONLoader
	operatorAPISchemasRegistry map[string]map[string]interface{}
}

func NewValidator() (*Client, error) {
	eventSchemas, eventSchemaRegistry, err := loadSchemas(pathSchemasRealTime, notificationTypes)
	if err != nil {
		return nil, err
	}

	operatorAPISchemas, operatorAPISchemasRegistry, err := loadSchemas(pathSchemasOperatorAPI, operatorAPIEndpoints)
	if err != nil {
		return nil, err
	}

	return &Client{
		eventSchemas:        eventSchemas,
		eventSchemaRegistry: eventSchemaRegistry,

		operatorAPISchemas:         operatorAPISchemas,
		operatorAPISchemasRegistry: operatorAPISchemasRegistry,
	}, nil
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetEventSchemas() map[string]map[string]interface{} {
	return c.eventSchemaRegistry
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetOperatorAPISchemas() map[string]map[string]interface{} {
	return c.operatorAPISchemasRegistry
}
