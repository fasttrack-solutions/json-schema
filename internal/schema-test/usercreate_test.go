package schemas

import (
	"path/filepath"
	"testing"
)

var pathUserCreateV2 = filepath.Join(realtTimeSchemaPath, "/user_create_v2.schema.json")

func TestValidData_Realtime_Post_UserRegistration(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "7865312321",
				"note": "Any note",
				"user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.106 Safari/537.36",
				"ip_address": "1.1.1.1",
				"timestamp": "2015-03-02T8:27:58.721607Z",
				"origin": "sub.example.com"
			}`,
			validTest: true,
		},
	}

	runTestCases(t, pathUserCreateV2, tests)
}
