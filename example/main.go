package main

import (
	"fmt"
	"os"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

func main() {

	validatorClient, err := validator.NewValidator()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Will validate a real time event
	realtimeEventTestPayload := `{
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

	notificationType := "payment"
	resultValidateEvent, errValidateEvent := validatorClient.ValidateEvent(notificationType, []byte(realtimeEventTestPayload))
	if errValidateEvent != nil {
		fmt.Printf("failed to validate json document: %v\n", errValidateEvent)
		os.Exit(1)
	}

	if len(resultValidateEvent) == 0 {
		fmt.Printf("real event %s validated, no errors found\n", notificationType)
	}

	if len(resultValidateEvent) > 0 {
		fmt.Printf("real event %s validated, found %d errors:\n", notificationType, len(resultValidateEvent))
		for _, e := range resultValidateEvent {
			fmt.Printf("path: %s. error: %s.\n", e.Path, e.Error)
		}
	}

	// Will validate an operator API response
	operatorAPITestResponse := `{
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

	endpoint := "user_details"
	resultValidateOperatorAPIResponse, errValidateOperatorAPIResponse := validatorClient.ValidateOperatorAPIResponse(endpoint, []byte(operatorAPITestResponse))
	if errValidateOperatorAPIResponse != nil {
		fmt.Printf("failed to validate json document: %v\n", errValidateOperatorAPIResponse)
		os.Exit(1)
	}

	if len(resultValidateOperatorAPIResponse) == 0 {
		fmt.Printf("operator api endpoint %s response validated, no errors found\n", endpoint)
	}

	if len(resultValidateOperatorAPIResponse) > 0 {
		fmt.Printf("operator api endpoint %s response validated, found %d errors:\n", endpoint, len(resultValidateOperatorAPIResponse))
		for _, e := range resultValidateOperatorAPIResponse {
			fmt.Printf("path: %s. error: %s.\n", e.Path, e.Error)
		}
	}
}
