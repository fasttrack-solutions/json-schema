package schemas

import (
	"path/filepath"
	"testing"
)

var customSMSSchema = filepath.Join(operatorApiSchemaPath, "/post/custom_sms.schema.json")

func TestValidData_Realtime_Post_SMS(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity": {
					"id": "12312-12312-123123-123123",
					"brand_id": 1,
					"user_id": "1123111",
					"internal_user_id": 1123112,
					"activity_id": 123,
					"action_group_id": 443,
					"trigger_hash": "asd777dsDSD778sd80-asdas",
					"action_id": 1224
				},
				"content": "TEST Deposit Free Bet 22 - 13th April",
				"recipient": "254716595289",
				"sender_name": "Casino1",
				"provider_name": "LinkMobility",
				"enconding": "auto"
			}`,
			isValid: true,
		},
	}

	runTestCases(t, customSMSSchema, tests, nil)
}
