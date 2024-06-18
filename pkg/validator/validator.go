package validator

import (
	"github.com/xeipuuv/gojsonschema"
)

type IValidator interface {
	GetEventSchemas() map[string]map[string]interface{}
	GetOperatorAPISchemas() map[string]map[string]interface{}
	ValidateRealTimeEvent(notificationType string, payload []byte) ([]ValidationError, error)
	ValidateOperatorAPIResponse(endpoint string, payload []byte) ([]ValidationError, error)
}

type ValidationError struct {
	Path  string
	Error string
}

type ValidationClient struct {
	realTimeSchemas, operatorAPISchemas                map[string]gojsonschema.JSONLoader
	realTimeSchemaRegistry, operatorAPISchemasRegistry map[string]map[string]interface{}
}

func NewValidator() (*ValidationClient, error) {
	eventSchemas, eventSchemaRegistry, err := loadSchemas(pathSchemasRealTime, getNotificationTypes())
	if err != nil {
		return nil, err
	}

	operatorAPISchemas, operatorAPISchemasRegistry, err := loadSchemas(pathSchemasOperatorAPI, getOperatorAPIEndpoints())
	if err != nil {
		return nil, err
	}

	return &ValidationClient{
		realTimeSchemas:            eventSchemas,
		realTimeSchemaRegistry:     eventSchemaRegistry,
		operatorAPISchemas:         operatorAPISchemas,
		operatorAPISchemasRegistry: operatorAPISchemasRegistry,
	}, nil
}

// GetEventSchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *ValidationClient) GetEventSchemas() map[string]map[string]interface{} {
	return c.realTimeSchemaRegistry
}

// GetOperatorAPISchemas returns schemas for real time events in a map string (payload type) gojsonschema
func (c *ValidationClient) GetOperatorAPISchemas() map[string]map[string]interface{} {
	return c.operatorAPISchemasRegistry
}

func getNotificationTypes() []string {
	return []string{
		"bonus",
		"cart",
		"casino",
		"custom",
		"game",
		"login_v2",
		"lottery_v2",
		"lottery",
		"payment",
		"sportsbook",
		"user_balances_update",
		"user_block_v2",
		"user_consents_v2",
		"user_create_v2",
		"user_update_v2",
	}
}

func getOperatorAPIEndpoints() []string {
	return []string{
		"user_details",
		"user_blocks",
		"user_consents",
	}
}
