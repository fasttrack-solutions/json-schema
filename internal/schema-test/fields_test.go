package schemas

import (
	"encoding/json"
	"fmt"
)

type fieldValidator struct {
	fieldName string
	validator func(string, schemaTest, *EventEnums) ([]schemaTest, error)
}

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

var fieldValidators = []fieldValidator{
	{keyAmount, validateAmountField},
	{keyBonusCode, validateBonusCodeField},
	{keyBonusID, validateBonusIDField},
	{keyCurrency, validateCurrencyField},
	{keyExchangeRate, validateExchangeRateField},
	{keyGameID, validateGameIDField},
	{keyLockedAmount, validateLockedAmountField},
	{keyProduct, validateProductField},
	{keyStatus, validateStatusField},
	{keyTimestamp, validateTimestampField},
	{keyType, validateTypeField},
	{keyUserBonusID, validateUserBonusIDField},
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
			name:    fmt.Sprintf("%s-%s", fieldName, testCase.name),
			payload: string(modifiedDocument),
			isValid: testCase.valid,
		}
	}

	return modifiedTestCases, nil
}

func validateAmountField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyAmount, schema, loadFieldTestsFloats(&FloatSettings{
		allowNegatives: false,
		allowZero:      true,
		allowNull:      false,
	}))
}

func validateBonusCodeField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyBonusCode, schema, loadFieldTestsString("BonusCode", nil))
}

func validateBonusIDField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyBonusID, schema, loadFieldTestsString("Bonus", &StringSettings{
		allowEmpty:     false,
		allowLowerCase: true,
		allowNumbers:   true,
		allowSpace:     true,
		allowSpecial:   true,
		allowUpperCase: true,
	}))
}

func validateCurrencyField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyCurrency, schema, loadFieldTestsString("EUR", &StringSettings{
		allowEmpty:     false,
		allowLowerCase: false,
		allowNumbers:   false,
		allowSpace:     false,
		allowSpecial:   false,
		allowUpperCase: true,
		includeNumeric: true,
	}))
}

func validateExchangeRateField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyExchangeRate, schema, loadFieldTestsFloats(&FloatSettings{
		allowNegatives: false,
		allowZero:      true,
		allowNull:      false,
	}))
}

func validateGameIDField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyGameID, schema, loadFieldTestsString("Origin", &StringSettings{
		allowEmpty:     true,
		allowLowerCase: true,
		allowNumbers:   true,
		allowSpace:     true,
		allowSpecial:   true,
		allowUpperCase: true,
		includeNumeric: true,
	}))
}

func validateLockedAmountField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyLockedAmount, schema, loadFieldTestsFloats(&FloatSettings{
		allowNegatives: false,
		allowZero:      true,
		allowNull:      false,
	}))
}

func validateOriginField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyOrigin, schema, loadFieldTestsString("Origin", &StringSettings{
		allowEmpty:     true,
		allowLowerCase: true,
		allowNumbers:   true,
		allowSpace:     false,
		allowSpecial:   true,
		allowUpperCase: true,
	}))
}

func validateProductField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyProduct, schema, loadFieldTestsString("Casino", &StringSettings{
		allowEmpty:     false,
		allowLowerCase: false,
		allowNumbers:   false,
		allowSpace:     false,
		allowSpecial:   false,
		allowUpperCase: false,
	}))
}

func validateStatusField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyStatus, schema, loadFieldTestsByEnums("Approved", enums.statusEnums, &StringSettings{
		allowEmpty:     false,
		allowLowerCase: false,
		allowNumbers:   false,
		allowSpace:     false,
		allowSpecial:   false,
		allowUpperCase: false,
	}))
}

func validateTypeField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyType, schema, loadFieldTestsByEnums("Bet", enums.typeEnums, &StringSettings{
		allowEmpty:     false,
		allowLowerCase: false,
		allowNumbers:   false,
		allowSpace:     false,
		allowSpecial:   false,
		allowUpperCase: false,
	}))
}

func validateUserBonusIDField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyUserBonusID, schema, loadFieldTestsString("Bonus", &StringSettings{
		allowEmpty:     false,
		allowLowerCase: true,
		allowNumbers:   true,
		allowSpace:     true,
		allowSpecial:   true,
		allowUpperCase: true,
	}))
}

func validateUserIDField(schemaPath string, schema schemaTest, enums *EventEnums) ([]schemaTest, error) {
	return addFieldTestCases(keyUserID, schema, loadFieldTestsString("12345", &StringSettings{
		allowEmpty:     false,
		allowLowerCase: true,
		allowNumbers:   true,
		allowSpace:     false,
		allowSpecial:   false,
		allowUpperCase: true,
	}))
}

func loadFieldTestsByEnums(defaultInput string, enumArray []string, settings *StringSettings) []validationCase {
	var testCases []validationCase
	if len(enumArray) == 0 || enumArray == nil {
		return loadFieldTestsString(defaultInput, settings)
	}

	for _, value := range enumArray {
		testCases = append(testCases, loadFieldTestsString(value, settings)...)
	}

	return testCases
}
