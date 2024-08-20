package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
	"github.com/stretchr/testify/require"
)

const (
	operatorApiSchemaPath = "../../pkg/validator/schemas/operator-api"
	realtTimeSchemaPath   = "../../pkg/validator/schemas/real-time-events"
)

type schemaTest struct {
	name    string
	payload string
	isValid bool
}

type validationCase struct {
	name  string
	value interface{}
	valid bool
}

type EventEnums struct {
	statusEnums []string
	typeEnums   []string
}

func runRealtimeTest(t *testing.T, schemaPath string, testCases []schemaTest, enums *EventEnums) {
	runTestCases(t, schemaPath, "real-time", testCases, enums)
}

func runOperatorTest(t *testing.T, schemaPath string, testCases []schemaTest, enums *EventEnums) {
	runTestCases(t, schemaPath, "operator-api", testCases, enums)
}

func runTestCases(t *testing.T, schemaPath, operatorType string, testCases []schemaTest, enums *EventEnums) {
	fieldTestCases, err := getFieldTests(schemaPath, testCases[0], enums)
	require.NoError(t, err)

	testCases = append(testCases, fieldTestCases...)

	for _, tc := range testCases {
		testName := generateTestName(tc.name, tc.isValid)

		t.Run(testName, func(t *testing.T) {
			client, err := validator.NewClient(validator.ValidationConfig{
				ChangePaymentNegativeCountToZero: false,
				DisableCurrencyValidation:        true,
				EnableStringIDs:                  false,
			})
			require.NoError(t, err)

			result := []validator.ValidationError{}
			switch operatorType {
			case "real-time":
				result, err = client.ValidateRealTimeEvent(schemaPath, []byte(tc.payload))
			case "operator-api":
				result, err = client.ValidateOperatorAPIResponse(schemaPath, []byte(tc.payload))
			}
			require.NoError(t, err)

			if !tc.isValid {
				require.Greater(t, len(result), 0, "Should have more than one error: %s, Result: %+v", tc, result)
				return
			}

			require.Len(t, result, 0, "Should have no errors: %s, Result: %+v", tc, result)
		})
	}
}

func getFieldTests(schema string, test schemaTest, enums *EventEnums) ([]schemaTest, error) {
	var (
		jsonData  map[string]interface{}
		testCases []schemaTest
		errors    []error
	)

	methodMap := make(map[string]func(string, schemaTest, *EventEnums) ([]schemaTest, error))
	for _, validator := range fieldValidators {
		methodMap[validator.fieldName] = validator.validator
	}

	err := json.Unmarshal([]byte(test.payload), &jsonData)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal payload: %v", err)
	}

	if enums == nil {
		enums = &EventEnums{}
	}

	for key := range jsonData {
		method, ok := methodMap[key]
		if ok {
			cassefs, err := method(schema, test, enums)
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

func generateTestName(name string, valid bool) string {
	validString := "Valid"
	if !valid {
		validString = "Invalid"
	}

	return fmt.Sprintf("%s_%s", name, strings.ToUpper(validString))
}
