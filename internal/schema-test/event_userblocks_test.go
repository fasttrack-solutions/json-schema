package schemas

import (
	"path/filepath"
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

var (
	pathUserBlockV2 = filepath.Join(realtTimeSchemaPath, "/user_block_v2.schema.json")
)

func TestValidData_Operator_Get_UserBlocks(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"blocks": [
					{
						"active": true,
						"type": "Excluded",
						"note": "Exclusion reason"
					},
					{
						"active": true,
						"type": "Blocked",
						"note": "Block reason"
					}
				]
			}`,
			isValid: true,
		},
	}

	runOperatorTest(t, validator.GetUserBlocks, tests, nil)
}

func TestValidData_Realtime_Put_UserBlock(t *testing.T) {
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

	runRealtimeTest(t, validator.NotificationTypeUserBlockV2, tests, nil)
}
