package schemas

func loadFieldTestsFloats() []validationCase {
	return []validationCase{
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
		{
			name:  "Null input",
			value: nil,
			valid: false,
		},
	}
}
