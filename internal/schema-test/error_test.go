package schemas

import (
	"path/filepath"
	"testing"
)

var errorSchema = filepath.Join(operatorApiSchemaPath, "/error_handling.schema.json")

func TestValidData_Operator_Error(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid error response",
			payload: `{
				"error": "invalid_request",
				"code": 400
			}`,
			isValid: true,
		},
	}

	runTestCases(t, errorSchema, tests, nil)
}
