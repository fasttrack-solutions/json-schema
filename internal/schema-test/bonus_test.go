package schemas

import (
	"path/filepath"
	"testing"
)

var (
	bonusListSchema        = filepath.Join(operatorApiSchemaPath, "/get/bonus_list.schema.json")
	bonusCreditSchema      = filepath.Join(operatorApiSchemaPath, "/post/bonus_credit.schema.json")
	bonusCreditFundsSchema = filepath.Join(operatorApiSchemaPath, "/post/post_bonus_credit_funds.schema.json")
	bonusRealtimeSchema    = filepath.Join(realtTimeSchemaPath, "/bonus.schema.json")
)

func TestValidData_Operator_Get_BonusList(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
						"Data": [{"text": "Bonus 1", "value": "123"}, {"text": "Bonus 2", "value": "456"}],
						"Success": true,
						"Errors": []
					}`,
			failsValidation: false,
		},
		{
			name: "Invalid data",
			payload: `{
				"Data": [{"text": 123, "value": "B1"}],
				"Success": true,
				"Errors": []
			}`,
			failsValidation: true,
		},
	}

	runTests(t, bonusListSchema, tests)
}

func TestValidData_Operator_Post_BonsuCredit(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "123abc",
				"bonus_code": "xyz"
			}`,
			failsValidation: false,
		},
	}

	runTests(t, bonusCreditSchema, tests)
}

func TestValidData_Operator_Post_BonusCreditFunds(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "123abc",
				"bonus_code": "xyz",
				"amount":  10.0,
				"currency": "USD"
			}`,
			failsValidation: false,
		},
	}

	runTests(t, bonusCreditFundsSchema, tests)
}

func TestValidData_Realtime_Post_Bonus(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"amount": 10.3,
				"bonus_code": "WELCOME100",
				"bonus_id": "9821",
				"bonus_turned_real": 0,
				"currency": "EUR",
				"exchange_rate": 0.1,
				"locked_amount": 2.3,
				"meta":{
					"key1": 10.00,
					"key2": "some string",
					"key3": false
				},
				"origin": "sub.example.com",
				"product": "Casino",
				"required_wagering_amount": 412.3,
				"status": "Created",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "WelcomeBonus",
				"user_bonus_id": "863512",
				"user_id": "7865312321"
			}`,
			failsValidation: false,
		},
	}

	runTests(t, bonusRealtimeSchema, tests)
}
