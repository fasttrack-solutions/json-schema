package main

import (
	"fmt"
	"os"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

func main() {
	result, err := validator.ValidateJSON("payment", []byte(testPayload))
	if err != nil {
		fmt.Printf("failed to validate json document: %v\n", err)
		os.Exit(1)
	}

	if len(result) == 0 {
		fmt.Println("json document validated, no errors found")
	}

	if len(result) > 0 {
		fmt.Printf("json document validated, found %d errors:\n", len(result))
		for _, e := range result {
			fmt.Println(e)
		}
	}
}

const testPayload = `{
  "amount": "32.76",
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
}`
