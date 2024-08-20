package benchmark

import (
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

// BenchmarkValidateRealTimeEvent benchmarks the loadSchemas function.
func BenchmarkValidateRealTimeEvent(b *testing.B) {
	realtimeEventTestPayload := `{
		"amount": 51.18,
		"bonus_code": "BONUSCODE",
		"currency": "USD",
		"exchange_rate": 0.91,
		"fee_amount": 1.61,
		"note": "string",
		"origin": "sub.example.com",
		"payment_id": "123456",
		"status": "Approved",
		"timestamp": "2015-03-02T8:27:58.721607Z",
		"type": "Debit",
		"user_id": 123457,
		"vendor_id": "123",
		"vendor_name": "Vendor"
	}`

	client, err := validator.NewClient(validator.ValidationConfig{})
	if err != nil {
		b.Fatalf("failed to create client: %v", err)
	}

	for i := 0; i < b.N; i++ {
		client.ValidateRealTimeEvent("PAYMENT", []byte(realtimeEventTestPayload))
	}
}
