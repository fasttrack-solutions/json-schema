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
	fieldTestCases, err := validateFields(schemaPath, testCases[0])
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

func generateTestName(name string, valid bool) string {
	validString := "Valid"
	if !valid {
		validString = "Invalid"
	}

	return fmt.Sprintf("%s_%s", name, strings.ToUpper(validString))
}
