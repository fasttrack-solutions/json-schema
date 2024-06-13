package schemas

import (
	"path/filepath"
	"testing"
)

var paymentSchema = filepath.Join(realtTimeSchemaPath, "/payment.schema.json")

func TestValidData_Realtime_Payment(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"amount": 32.76,
				"bonus_code": "CHRISTMAS2023",
				"currency": "USD",
				"exchange_rate": 0.91,
				"fee_amount": 2.34,
				"note": "string",
				"origin": "sub.example.com",
				"payment_id": "23541",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607Z",
				"type": "Debit",
				"user_id": "7865312321",
				"vendor_id": "562",
				"vendor_name": "Skrill"
			}`,
			isValid: true,
		},
	}

	runTestCases(t, paymentSchema, tests, &EventEnums{
		statusEnums: []string{"Approved"},
		typeEnums:   []string{"Debit"},
	})
}
