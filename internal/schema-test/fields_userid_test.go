package schemas

import (
	"testing"
)

func validateUserIDField(t *testing.T, schemaPath string, test schemaTest) {
	var userTestCases = []validationCase{
		{
			name:            "Input as string",
			value:           "123",
			failsValidation: false,
		},
		{
			name:            "Input as numeric",
			value:           123,
			failsValidation: true,
		},
	}

	validateField(t, "User ID", "user_id", schemaPath, test, userTestCases)
}
