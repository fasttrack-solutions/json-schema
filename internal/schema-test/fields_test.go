package schemas

import (
	"encoding/json"
	"fmt"
)

const (
	keyAmount       = "amount"
	keyBonusCode    = "bonus_code"
	keyBonusID      = "bonus_id"
	keyCurrency     = "currency"
	keyExchangeRate = "exchange_rate"
	keyGameID       = "game_id"
	keyLockedAmount = "locked_amount"
	keyOrigin       = "origin"
	keyProduct      = "product"
	keyStatus       = "status"
	keyTimestamp    = "timestamp"
	keyType         = "type"
	keyUserBonusID  = "user_bonus_id"
	keyUserID       = "user_id"
)

type fieldValidator struct {
	fieldName string
	validator func(string, schemaTest) ([]schemaTest, error)
}

var fieldValidators = []fieldValidator{
	{keyAmount, validateAmountField},
	{keyBonusCode, validateBonusCodeField},
	{keyBonusID, validateBonusIDField},
	{keyCurrency, validateCurrencyField},
	{keyExchangeRate, validateExchangeRateField},
	{keyGameID, validateGameIDField},
	{keyLockedAmount, validateLockedAmountField},
	{keyOrigin, validateOriginField},
	{keyProduct, validateProductField},
	{keyStatus, validateStatusField},
	{keyTimestamp, validateTimestampField},
	{keyType, validateTypeField},
	{keyUserBonusID, validateUserBonusIDField},
	{keyUserID, validateUserIDField},
}

func validateFields(schema string, test schemaTest) ([]schemaTest, error) {
	var (
		jsonData  map[string]interface{}
		testCases []schemaTest
		errors    []error
	)

	methodMap := make(map[string]func(string, schemaTest) ([]schemaTest, error))
	for _, validator := range fieldValidators {
		methodMap[validator.fieldName] = validator.validator
	}

	err := json.Unmarshal([]byte(test.payload), &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal payload: %v", err)
	}

	for key := range jsonData {
		method, ok := methodMap[key]
		if ok {
			cassefs, err := method(schema, test)
			if err != nil {
				errors = append(errors, fmt.Errorf("Failed to validate field %s: %v", key, err))
				continue
			}

			testCases = append(testCases, cassefs...)
		}
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("Failed to validate payload: %v", errors)
	}

	return testCases, nil
}

func addFieldTestCases(fieldName string, test schemaTest, testCases []validationCase) ([]schemaTest, error) {
	modifiedTestCases := make([]schemaTest, len(testCases))

	for i, testCase := range testCases {
		var payload map[string]interface{}
		err := json.Unmarshal([]byte(test.payload), &payload)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal payload: %v", err)
		}

		_, ok := payload[fieldName]
		if !ok {
			return nil, fmt.Errorf("%s not found in payload: %s", fieldName, payload)
		}

		payload[fieldName] = testCase.value

		modifiedDocument, err := json.Marshal(payload)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal payload: %v", err)
		}

		modifiedTestCases[i] = schemaTest{
			name:      fmt.Sprintf("%s-%s", fieldName, testCase.name),
			payload:   string(modifiedDocument),
			validTest: testCase.valid,
		}
	}

	return modifiedTestCases, nil
}

func validateAmountField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyAmount, test, loadFieldTestsFloats())
}

func validateBonusCodeField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyBonusCode, test, loadFieldTestsString("Bonus", false))
}

func validateBonusIDField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyBonusID, test, loadFieldTestsString("Bonus", false))
}

func validateCurrencyField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyCurrency, test, loadFieldTestsString("EUR", false))
}

func validateExchangeRateField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyExchangeRate, test, loadFieldTestsString("Bonus", false))
}

func validateGameIDField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyGameID, test, loadFieldTestsString("Origin", false))
}

func validateLockedAmountField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyLockedAmount, test, loadFieldTestsString("Bonus", false))
}

func validateOriginField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyOrigin, test, loadFieldTestsString("Origin", false))
}

func validateProductField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyProduct, test, loadFieldTestsString("Bonus", false))
}

func validateStatusField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyStatus, test, loadFieldTestsString("Pending", false))
}

func validateTypeField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyType, test, loadFieldTestsString("Bonus", false))
}

func validateUserBonusIDField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyUserBonusID, test, loadFieldTestsString("Bonus", false))
}

func validateUserIDField(schemaPath string, test schemaTest) ([]schemaTest, error) {
	return addFieldTestCases(keyUserID, test, loadFieldTestsString("UserID", false))
}
