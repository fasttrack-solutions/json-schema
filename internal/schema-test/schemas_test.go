package schemas

import (
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

func runTestCases(t *testing.T, schemaPath string, testCases []schemaTest, enums *EventEnums) {
	fieldTestCases, err := createFieldTests(schemaPath, testCases[0], enums)
	require.NoError(t, err)

	testCases = append(testCases, fieldTestCases...)

	for _, tc := range testCases {
		testName := generateTestName(tc.name, tc.isValid)

		t.Run(testName, func(t *testing.T) {
			schemaLoader := gojsonschema.NewReferenceLoader("file://" + schemaPath)
			documentLoader := gojsonschema.NewStringLoader(tc.payload)

			result, err := gojsonschema.Validate(schemaLoader, documentLoader)
			require.NoError(t, err)

			if !tc.isValid {
				require.False(t, result.Valid(), "Should have errors: %s, Result: %+v", tc, result.Errors())
				return
			}

			require.True(t, result.Valid(), "Should have no errors: %s, Result: %+v", tc, result.Errors())
		})
	}
}

func generateTestName(name string, valid bool) string {
	validString := "Valid"
	if !valid {
		validString = "Invalid"
	}

	return fmt.Sprintf("%s_%s", name, strings.ToUpper(validString))
}
