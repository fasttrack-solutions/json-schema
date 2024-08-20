package schemas

import (
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

func TestValidData_Realtime_Post_UserBalances(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"balances": [
					{
						"amount": 50,
						"currency": "EUR",
						"key": "real_money",
						"exchange_rate": 1
					},
					{
						"amount": 50,
						"currency": "EUR",
						"key": "bonus_money",
						"exchange_rate": 1
					}
				],
				"origin": "example.com",
				"timestamp": "2019-07-16T12:00:00Z",
				"user_id": "1234"
			}`,
			isValid: true,
		},
	}

	runRealtimeTest(t, validator.NotificationTypeUserBalancesUpdate, tests, nil)
}
