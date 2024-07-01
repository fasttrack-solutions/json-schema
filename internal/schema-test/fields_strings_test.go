package schemas

import (
	"fmt"
	"strings"
)

type StringSettings struct {
	allowEmpty     bool
	allowLowerCase bool
	allowNumbers   bool
	allowSpace     bool
	allowSpecial   bool
	allowUpperCase bool
}

func createTestCases(valid bool, testCases map[string]interface{}) []validationCase {
	result := make([]validationCase, 0, len(testCases))

	for testName, testString := range testCases {
		result = append(result, validationCase{
			name:  testName,
			value: testString,
			valid: valid,
		})
	}

	return result
}

func loadFieldTestsString(input string, config *StringSettings) []validationCase {
	if config == nil {
		config = &StringSettings{
			allowEmpty:     true,
			allowLowerCase: true,
			allowNumbers:   true,
			allowSpace:     true,
			allowSpecial:   true,
			allowUpperCase: true,
		}
	}

	defaultTests := createTestCases(true, map[string]interface{}{
		"String input": input,
	})

	numberTests := createTestCases(config.allowNumbers, map[string]interface{}{
		"String input as numbers":   "123",
		"String input with numbers": fmt.Sprintf("%s123", input),
	})

	lowerCaseTests := createTestCases(config.allowLowerCase, map[string]interface{}{
		"String input to lower case": strings.ToLower(input),
	})

	upperCaseTests := createTestCases(config.allowUpperCase, map[string]interface{}{
		"String input to upper case": strings.ToUpper(input),
	})

	spacingTests := createTestCases(config.allowSpace, map[string]interface{}{
		"String input with leading and trailing spaces": fmt.Sprintf(" %s ", input),
		"String input with spaces":                      fmt.Sprintf("test %s 123", input),
	})

	specialCharTests := createTestCases(config.allowSpecial, map[string]interface{}{
		"String input with special characters and spaces": fmt.Sprintf("test@%s#123 £!$^&*()", input),
	})

	emptyStringTests := createTestCases(config.allowEmpty, map[string]interface{}{
		"Empty input": "",
	})

	nonStringTests := []validationCase{
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

	cases := make([]validationCase, 0)
	cases = append(cases, defaultTests...)
	cases = append(cases, numberTests...)
	cases = append(cases, lowerCaseTests...)
	cases = append(cases, upperCaseTests...)
	cases = append(cases, spacingTests...)
	cases = append(cases, specialCharTests...)
	cases = append(cases, emptyStringTests...)
	cases = append(cases, nonStringTests...)

	return cases
}
