package schemas

func validateTimestampField(schemaPath string, test schemaTest, enums *EventEnums) ([]schemaTest, error) {
	var timestampTestCases = []validationCase{
		{
			name:  "Valid format",
			value: "2015-03-02T8:27:58.721607Z",
			valid: true,
		},
		{
			name:  "Valid format with offset",
			value: "2015-03-02T8:27:58.721607+01:00",
			valid: true,
		},
		{
			name:  "Valid format with missing milliseconds",
			value: "2015-03-02T8:27.721607Z",
			valid: false,
		},
		{
			name:  "Valid format without Z",
			value: "2015-03-02T8:27.721607Z",
			valid: false,
		},
		{
			name:  "Invalid format",
			value: "2015-03",
			valid: false,
		},
		{
			name:  "Timestamp as string of numbers",
			value: "20150302",
			valid: false,
		},
		{
			name:  "Numeric format",
			value: 20150302,
			valid: false,
		},
		{
			name:  "Boolean input",
			value: false,
			valid: false,
		},
		{
			name:  "Null input",
			value: nil,
			valid: false,
		},
	}

	return addFieldTestCases("timestamp", test, timestampTestCases)
}
