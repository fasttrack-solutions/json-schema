package schemas

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

const (
	operatorApiSchemaPath = "../../pkg/validator/schemas/operator-api"
	realtTimeSchemaPath   = "../../pkg/validator/schemas/real-time-events"
)

type schemaTest struct {
	name      string
	payload   string
	validTest bool
}

type validationCase struct {
	name  string
	value interface{}
	valid bool
}

func runTestCases(t *testing.T, schemaPath string, testCases []schemaTest) {
	fieldTestCases, err := validatePayload(t, schemaPath, testCases[0])
	require.NoError(t, err)

	testCases = append(testCases, fieldTestCases...)

	for _, tc := range testCases {
		testName := generateTestName(tc.name, tc.validTest)

		t.Run(testName, func(t *testing.T) {
			schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
			documentLoader := gojsonschema.NewStringLoader(tc.payload)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			require.NoError(t, err)

			if !tc.validTest {
				require.False(t, result.Valid(), "The document should not be valid %s, %v", tc, result.Errors())
				return
			}

			require.True(t, result.Valid(), "The document should be valid", result.Errors())
		})
	}
}

func validateField(t *testing.T, testName, fieldName, schemaPath string, test schemaTest, testCases []validationCase) []schemaTest {
	modifiedTestCases := make([]schemaTest, len(testCases))

	for i, testCase := range testCases {
		var payload map[string]interface{}
		err := json.Unmarshal([]byte(test.payload), &payload)
		if err != nil {
			t.Fatalf("Failed to unmarshal payload: %v", err)
		}

		_, ok := payload[fieldName]
		if !ok {
			t.Fatalf("%s not found in payload: %s", fieldName, payload)
		}

		payload[fieldName] = testCase.value

		modifiedDocument, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("failed to marshal payload: %v", err)
		}

		modifiedTestCases[i] = schemaTest{
			name:      testName + " " + testCase.name,
			payload:   string(modifiedDocument),
			validTest: testCase.valid,
		}
	}

	return modifiedTestCases
}

func generateTestName(name string, valid bool) string {
	validString := "Valid"
	if !valid {
		validString = "Invalid"
	}

	return fmt.Sprintf("%s_%s", name, strings.ToUpper(validString))
}

func validatePayload(t *testing.T, schema string, test schemaTest) ([]schemaTest, error) {
	var payloadstr map[string]interface{}
	var testCases []schemaTest

	methodMap := map[string]func(*testing.T, string, schemaTest) []schemaTest{
		keyCurrency: validateCurrencyField,
		keyGameID:   validateOriginField,
		keyOrigin:   validateStatusField,
		keyStatus:   validateTimestampField,
		keyUserID:   validateUserIDField,
	}

	err := json.Unmarshal([]byte(test.payload), &payloadstr)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal payload: %v", err)
	}

	for key := range payloadstr {
		method, ok := methodMap[key]
		if ok {
			testCases = append(testCases, method(t, schema, test)...)
		}
	}

	return testCases, nil
}
