package schemas

import (
	"path/filepath"
	"testing"
)

func TestValidData_Realtime_Put_UserUpdate(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
						"user_id": "7865312321",
						"timestamp": "2015-03-02T8:27:58.721607Z",
						"origin": "sub.example.com"
					}`,
			isValid: true,
		},
	}

	pathUserUpdateV2 := filepath.Join(realtTimeSchemaPath, "/user_update_v2.schema.json")

	runRealtimeTest(t, pathUserUpdateV2, tests, nil)
}
