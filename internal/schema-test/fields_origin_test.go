package schemas

import (
	"testing"
)

func validateOriginField(t *testing.T, schemaPath string, test schemaTest) {
	var originTestCases = []validationCase{
		{
			name:            "String input",
			value:           "testOrigin123",
			failsValidation: false,
		},
		{
			name:            "String input with spaces",
			value:           "test Origin 123",
			failsValidation: false,
		},
		{
			name:            "Empty string input",
			value:           "",
			failsValidation: false,
		},
		{
			name:            "String input with special characters",
			value:           "test@Origin#123",
			failsValidation: false,
		},
		{
			name:            "String input with only spaces",
			value:           "  ",
			failsValidation: true,
		},
		{
			name:            "String input with leading and trailing spaces",
			value:           " testOrigin123 ",
			failsValidation: true,
		},
		{
			name:            "Numeric input",
			value:           12345,
			failsValidation: true,
		},
		{
			name:            "Boolean input",
			value:           true,
			failsValidation: true,
		},
		{
			name:            "Null input",
			value:           nil,
			failsValidation: true,
		},
	}

	validateField(t, "Origin", "origin", schemaPath, test, originTestCases)
}
