package validator

import (
	"github.com/xeipuuv/gojsonschema"
)

var (
	notificationTypes = []string{
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

	operatorAPIEndpoints = []string{
		"user_details",
		"user_blocks",
		"user_consents",
	}
)

type Error struct {
	Path  string
	Error string
}

type Client struct {
	eventSchemas       map[string]gojsonschema.JSONLoader
	operatorAPISchemas map[string]gojsonschema.JSONLoader

	publicEventSchemas       map[string]map[string]interface{}
	publicOperatorAPISchemas map[string]map[string]interface{}
}

func NewValidator() (*Client, error) {
	eventSchemas, publicEventSchemas, err := loadSchemas(schemasPathRealTimeEvents, notificationTypes)
	if err != nil {
		return nil, err
	}

	operatorAPISchemas, publicOperatorAPISchemas, err := loadSchemas(schemasPathOperatorAPI, notificationTypes)
	if err != nil {
		return nil, err
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
