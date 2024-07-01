package schemas

type FloatSettings struct {
	allowNegatives bool
	allowZero      bool
	allowNull      bool
}

func loadFieldTestsFloats(config *FloatSettings) []validationCase {
	testCases := []validationCase{
		{
			name:  "Float input",
			value: 10.54,
			valid: true,
		},
		{
			name:  "Float input without decimal",
			value: 10,
			valid: true,
		},
		{
			name:  "Float input with 2 decimals",
			value: 10.54,
			valid: true,
		},
		{
			name:  "Float input with 5 decimals",
			value: 10.54334,
			valid: true,
		},
		{
			name:  "String input",
			value: "TestInput",
			valid: false,
		},
		{
			name:  "Boolean input",
			value: false,
			valid: false,
		},
	}

	negativeTestCases := createTestCases(config.allowNegatives, map[string]interface{}{
		"Negative float input": -10.54,
	})

	zeroTestCases := createTestCases(config.allowZero, map[string]interface{}{
		"Zero float input": 0,
	})

	nullTestCases := createTestCases(config.allowNull, map[string]interface{}{
		"Null input": nil,
	})

	testCases = append(testCases, zeroTestCases...)
	testCases = append(testCases, nullTestCases...)
	testCases = append(testCases, negativeTestCases...)

	return testCases
}
