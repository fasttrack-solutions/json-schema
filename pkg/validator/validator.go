package validator

import (
	"github.com/xeipuuv/gojsonschema"
)

type Client struct {
	realTimeSchemas, operatorAPISchemas                map[string]gojsonschema.JSONLoader
	realTimeSchemaRegistry, operatorAPISchemasRegistry map[string]map[string]interface{}
}

func NewClient() (*Client, error) {
	eventSchemas, eventSchemaRegistry, err := loadSchemas(pathSchemasRealTime, getNotificationTypes())
	if err != nil {
		return nil, err
	}

	operatorAPISchemas, operatorAPISchemasRegistry, err := loadSchemas(pathSchemasOperatorAPI, getOperatorAPIEndpoints())
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
		"user_tags",
		"user_migration",
		"poker_tournament_buy_in_v2",
		"poker_tournament_cash_out_v2",
		"poker_cash_game_buy_in_v2",
		"poker_cash_game_cash_out_v2",
	}
}

func getOperatorAPIEndpoints() []string {
	return []string{
		"user_details",
		"user_blocks",
		"user_consents",
	}
}
