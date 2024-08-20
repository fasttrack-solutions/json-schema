package validator

import (
	"github.com/xeipuuv/gojsonschema"
)

type Client struct {
	realTimeSchemas, operatorAPISchemas                map[string]gojsonschema.Schema
	realTimeSchemaRegistry, operatorAPISchemasRegistry map[string]map[string]interface{}
	ValidationConfig
}

func NewClient(config ValidationConfig) (*Client, error) {
	eventSchemas, eventSchemaRegistry, err := loadSchemas(pathSchemasRealTime, getNotificationTypes(), config)
	if err != nil {
		return nil, err
	}

	operatorAPISchemas, operatorAPISchemasRegistry, err := loadSchemas(pathSchemasOperatorAPI, getOperatorAPIEndpoints(), config)
	if err != nil {
		return nil, err
	}

	return &Client{
		realTimeSchemas:            eventSchemas,
		realTimeSchemaRegistry:     eventSchemaRegistry,
		operatorAPISchemas:         operatorAPISchemas,
		operatorAPISchemasRegistry: operatorAPISchemasRegistry,
	}, nil
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetEventSchemas() map[string]map[string]interface{} {
	return c.realTimeSchemaRegistry
}

// GetOperatorAPISchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *Client) GetOperatorAPISchemas() map[string]map[string]interface{} {
	return c.operatorAPISchemasRegistry
}
