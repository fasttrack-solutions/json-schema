package schemas

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

const (
	operatorApiSchemaPath = "../../pkg/validator/schemas/operator-api"
	realtTimeSchemaPath   = "../../pkg/validator/schemas/real-time-events"
)

type schemaTest struct {
	name            string
	payload         string
	failsValidation bool
}

type validationCase struct {
	name            string
	value           interface{}
	failsValidation bool
}

func runTests(t *testing.T, schemaPath string, tests []schemaTest) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
			documentLoader := gojsonschema.NewStringLoader(tt.payload)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			require.NoError(t, err)

			if tt.failsValidation {
				require.False(t, result.Valid(), "The document should not be valid")
				return
			}

			require.True(t, result.Valid(), "The document should be valid", result.Errors())
		})
	}
}

func validateField(t *testing.T, testName, fieldName, schemaPath string, test schemaTest, testCases []validationCase) {
	modifiedTestCases := make([]schemaTest, len(testCases))

	for i, testData := range testCases {
		var payload map[string]interface{}
		err := json.Unmarshal([]byte(test.payload), &payload)
		if err != nil {
			t.Fatalf("Failed to unmarshal document: %v", err)
		}

		_, ok := payload["user_id"].(string)
		if !ok {
			t.Fatal("user_id not found in document")
		}

		payload[fieldName] = testData.value

		modifiedDocument, err := json.Marshal(payload)
		if err != nil {
			t.Fatalf("failed to marshal document: %v", err)
		}

		modifiedTestCases[i] = schemaTest{
			name:            testName + " " + testData.name,
			payload:         string(modifiedDocument),
			failsValidation: testData.failsValidation,
		}
	}

	runTests(t, schemaPath, modifiedTestCases)
}
