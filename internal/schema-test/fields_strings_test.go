package schemas

import (
	"fmt"
	"strings"
)

const formatHostname = "hostname"

func loadFieldTestsString(input string, allowSpaceAndSpecials bool) []validationCase {
	cases := []validationCase{
		{
			name:  "String input",
			value: "%s",
			valid: true,
		},
		{
			name:  "String input with numbers",
			value: "%s123",
			valid: true,
		},
		{
			name:  "String input to lower case",
			value: strings.ToLower("%s"),
			valid: true,
		},
		{
			name:  "String input to upper case",
			value: strings.ToLower("%s"),
			valid: true,
		},
		{
			name:  "Numeric input",
			value: 12345,
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

	hostnameCases := []validationCase{
		{
			name:  "String input with spaces",
			value: "test %s 123",
		},
		{
			name:  "String input with special characters and spaces",
			value: "test@%s#123 £!$%^&*()",
		},
		{
			name:  "Empty string input",
			value: "",
		},
		{
			name:  "String input with only spaces",
			value: "  ",
		},
		{
			name:  "String input with leading and trailing spaces",
			value: " %s ",
		},
	}

	for i := range hostnameCases {
		hostnameCases[i].valid = allowSpaceAndSpecials
	}

	cases = append(cases, hostnameCases...)

	for i, c := range cases {
		if str, ok := c.value.(string); ok {
			cases[i].value = fmt.Sprintf(str, input)
		}
	}

	return cases
}
