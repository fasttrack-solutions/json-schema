package validator

import (
	"encoding/json"
	"fmt"
)

type ValidationConfig struct {
	ChangePaymentNegativeCountToZero bool
	DisableCurrencyValidation        bool
	EnableStringIDs                  bool
}

func processConfig(notificationType string, schema map[string]interface{}, config ValidationConfig) (map[string]interface{}, error) {
	notificationConfig := map[string][]bool{
		NotificationTypePayment: {
			config.ChangePaymentNegativeCountToZero,
			config.DisableCurrencyValidation,
			config.EnableStringIDs,
		},
		NotificationTypeLotteryV2: {
			config.DisableCurrencyValidation,
		},
		NotificationTypeBonus: {
			config.DisableCurrencyValidation,
		},
		NotificationTypeCasino: {
			config.EnableStringIDs,
		},
	}

	if processMap, exists := notificationConfig[notificationType]; exists {
		for _, action := range processMap {
			config.applyValidationConfig(action, notificationType, schema)
		}
	}

	return schema, nil
}

func (c *ValidationConfig) applyValidationConfig(setting bool, notificationType string, schema map[string]interface{}) (map[string]interface{}, error) {
	listProcess := map[bool]func(string, map[string]interface{}) (map[string]interface{}, error){
		c.ChangePaymentNegativeCountToZero: setChangePaymentNegativeCount,
		c.DisableCurrencyValidation:        setDisableCurrencyValidation,
		c.EnableStringIDs:                  setEnableStringIDs,
	}

	return listProcess[setting](notificationType, schema)
}

func replaceSchemaProperties(schema map[string]interface{}, replace map[string]string) (map[string]interface{}, error) {
	for key, value := range replace {
		var mapStr map[string]interface{}
		err := json.Unmarshal([]byte(value), &mapStr)
		if err != nil {
			return nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		if properties, ok := schema["properties"].(map[string]interface{}); ok {
			properties[key] = mapStr
		}
	}

	return schema, nil
}
