package validator

import (
	"embed"
	"errors"
	"fmt"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

const basePath = "validation_schemas_events"

//go:embed validation_schemas_events/*.json
var eventSchemasFS embed.FS

func ValidateJSON(notificationType string, payload []byte) ([]Error, error) {
	if notificationType == "" {
		return nil, errors.New("no notification type provided")
	}

	if payload == nil {
		return nil, errors.New("no payload provided")
	}

	filename := getSchemaForNotificationType(notificationType)

	if filename == "" {
		return nil, fmt.Errorf("no json schema filename found")
	}

	// Maybe we should use an instance of the validator instead of digging in its file system on each validation...
	schemaData, err := eventSchemasFS.ReadFile(basePath + "/" + filename)
	if err != nil {
		return nil, fmt.Errorf("error reading schema file: %v", err)
	}

	schemaLoader := gojsonschema.NewBytesLoader(schemaData)
	documentLoader := gojsonschema.NewBytesLoader(payload)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		return nil, err
	}

	if result.Valid() {
		return nil, nil
	}

	returnErrors := make([]Error, len(result.Errors()))
	for i, e := range result.Errors() {
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

func getSchemaForNotificationType(notificationType string) string {
	switch strings.ToLower(notificationType) {
	case "bonus":
		return "BONUS.schema.json"
	case "casino":
		return "CASINO.schema.json"
	case "custom":
		return "CUSTOM.schema.json"
	case "game":
		return "GAME.schema.json"
	case "login_v2":
		return "LOGIN_V2.schema.json"
	case "lottery_v2":
		return "LOTTERY_V2-purchase.schema.json"
	case "payment":
		return "PAYMENT.schema.json"
	case "sportsbook":
		return "SPORTSBOOK.schema.json"
	case "user_balances_update":
		return "USER_BALANCES_UPDATE.schema.json"
	case "user_block_v2":
		return "USER_BLOCK_V2.schema.json"
	case "user_consents_v2":
		return "USER_CONSENTS_V2.schema.json"
	case "user_create":
		return "USER_CREATE_V2.schema.json"
	case "user_update_v2":
		return "USER_UPDATE_V2.schema.json"
	default:
		return ""
	}
}
