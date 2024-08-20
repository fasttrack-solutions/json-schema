package schemas

import (
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

func TestValidData_Realtime_Post_CustomEvent(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"notification_type": "password_forgotten",
				"user_id": "123456",
				"user_agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_3) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.116 Safari/537.36",
				"ip_address": "192.168.1.1",
				"timestamp": "2020-03-02T8:27:58.721607+00:00",
				"origin": "sub.example.com",
				"data": {
					"password_unlock_code": "abc123"
				}
			}`,
			isValid: true,
		},
	}

	runRealtimeTest(t, validator.NotificationTypeCustom, tests, nil)
}
