package schemas

import (
	"path/filepath"
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

var loginSchema = filepath.Join(realtTimeSchemaPath, "/login_v2.schema.json")

func TestValidData_Realtime_Post_Login(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "7865312321",
				"is_impersonated": true,
				"ip_address": "192.0.2.1",
				"user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/68.0.3440.75 Safari/537.36",
				"timestamp": "2015-03-02T8:27:58.721607Z",
				"origin": "sub.example.com"
			}`,
			isValid: true,
		},
	}

	runRealtimeTest(t, validator.NotificationTypeLoginV2, tests, nil)
}
