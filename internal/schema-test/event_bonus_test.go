package schemas

import (
	"path/filepath"
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

var (
	bonusRealtimeSchema = filepath.Join(realtTimeSchemaPath, "/bonus.schema.json")
)

func TestValidData_OperatorGetBonusList(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"Data": [{"text": "Bonus 1", "value": "123"}, {"text": "Bonus 2", "value": "456"}],
				"Success": true,
				"Errors": []
			}`,
			isValid: true,
		},
		{
			name: "Invalid data",
			payload: `{
				"Data": [{"text": 123, "value": "B1"}],
				"Success": true,
				"Errors": []
			}`,
			isValid: false,
		},
	}

	runOperatorTest(t, validator.GetBonusList, tests, nil)
}

func TestValidData_OperatorPostBonsuCredit(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "123abc",
				"bonus_code": "xyz"
			}`,
			isValid: true,
		},
	}

	runOperatorTest(t, validator.PostBonusCredit, tests, nil)
}

func TestValidData_OperatorPostBonusCreditFunds(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"user_id": "123abc",
				"bonus_code": "xyz",
				"amount":  10.0,
				"currency": "USD"
			}`,
			isValid: true,
		},
	}

	runOperatorTest(t, validator.PostBonusCreditFunds, tests, nil)
}

func TestValidData_RealtimePostBonus(t *testing.T) {
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
			isValid: true,
		},
	}

	runRealtimeTest(t, validator.NotificationTypeBonus, tests, &EventEnums{
		statusEnums: []string{"Pending"},
		typeEnums:   []string{"WelcomeBonus"},
	})
}
