package validator

func setChangePaymentNegativeCount(notificationType string, schema map[string]interface{}) (map[string]interface{}, error) {
	const (
		depositCountStr = `{
				"description": "Number of deposits made by the user",
				"type": "number",
				"minimum": 0,
				"maximum": 0
			}`
		withdrawCountStr = `{
				"description": "Number of withdrawals made by the user",
				"type": "number",
				"minimum": 0,
				"maximum": 0
			}`
	)

	replacementMap := map[string]string{
		"deposit_count":  depositCountStr,
		"withdraw_count": withdrawCountStr,
	}

	return replaceSchemaProperties(schema, replacementMap)
}

func setDisableCurrencyValidation(notificationType string, schema map[string]interface{}) (map[string]interface{}, error) {
	const (
		currencyStr = `{
				"description": "Number of deposits made by the user",
				"type": "number",
				"minimum": 0,
				"maximum": 0
			}`
	)

	replacementMap := map[string]string{
		"currency": currencyStr,
	}

	return replaceSchemaProperties(schema, replacementMap)
}

func setEnableStringIDs(notificationType string, schema map[string]interface{}) (map[string]interface{}, error) {
	const (
		gameStr = `{
			"description": "Game ID",
			"type": "string"
		}`
		vendorStr = `{
			"description": "Vendor ID",
			"type": "number"
		}`
		paymentStr = `{
			"description": "Platform ID of the payment",
			"type": "number"
		}`
	)

	replacementMap := map[string]string{
		"game_id":    gameStr,
		"vendor_id":  vendorStr,
		"payment_id": paymentStr,
	}

	return replaceSchemaProperties(schema, replacementMap)
}
