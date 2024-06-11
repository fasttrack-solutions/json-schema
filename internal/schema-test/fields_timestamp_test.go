package schemas

import "testing"

func validateTimestampField(t *testing.T, schemaPath string, test schemaTest) {
	var timestampTestCases = []validationCase{
		{
			name:            "Valid format",
			value:           "2015-03-02T8:27:58.721607Z",
			failsValidation: false,
		},
		{
			name:            "Invalid format",
			value:           "2015-03",
			failsValidation: false,
		},
		{
			name:            "Numeric format",
			value:           20150302,
			failsValidation: true,
		},
	}

	validateField(t, "Timestamp", "timestamp", schemaPath, test, timestampTestCases)
}
