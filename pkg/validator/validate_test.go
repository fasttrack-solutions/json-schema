package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

func TestClient_ValidateRealTimeEvent(t *testing.T) {
	const (
		testValidEvent   = `{"amount":32.76,"bonus_code":"BONUS_CODE","currency":"USD","deposit_count":0,"exchange_rate":0.91,"fee_amount":2.34,"note":"string","origin":"test","payment_id":"1234","status":"approved","timestamp":"2023-01-01T07:29:58.721607Z","type":"credit","user_id":"12345","vendor_id":"123","vendor_name":"vendortest","withdraw_count":0}`
		testInvalidEvent = `{"amount":32.76,"bonus_code":"BONUS_CODE","currency":"USD","deposit_count":0,"exchange_rate":0.91,"fee_amount":2.34,"note":"string","origin":"test","payment_id":"1234","status":"approved","timestamp":"2023-01-01T07:29:58.721607Z","type":"credit","user_id":12345,"vendor_id":"123","vendor_name":"vendortest","withdraw_count":0}`
		schemaJson       = `{"$id":"https://www.fasttrack-solutions.com/en/resources/integration/real-time-data/login","$schema":"http://json-schema.org/draft-07/schema#","title":"LoginEvent(LOGIN_V2)","type":"object","required":["user_id","timestamp","origin"],"properties":{"user_id":{"type":"string","description":"Theuniqueidoftheuser"},"is_impersonated":{"type":"boolean","description":"Thisfieldistruewhenasupportagentisloggedinimpersonatingauser"},"ip_address":{"type":"string","description":"IPAddress"},"timestamp":{"type":"string","format":"date-time","description":"TimestampoflogininRFC3339format"},"origin":{"type":"string","description":"TheOriginoftheuser"}}}`
	)

	type args struct {
		notificationType string
		payload          []byte
	}

	tests := []struct {
		name                     string
		args                     args
		expectedValidationErrors []ValidationError
		expectError              bool
	}{
		{
			name: "Valid payload",
			args: args{
				notificationType: "login_v2",
				payload:          []byte(testValidEvent),
			},
			expectedValidationErrors: nil,
			expectError:              false,
		},
		{
			name: "Invalid payload",
			args: args{
				notificationType: "login_v2",
				payload:          []byte(testInvalidEvent),
			},
			expectedValidationErrors: []ValidationError{
				{
					Path:  "(root).user_id",
					Error: "Invalid type. Expected: string, given: integer",
				},
			},
			expectError: false,
		},
		{
			name: "Invalid notification type",
			args: args{
				notificationType: "invalid",
				payload:          []byte(testValidEvent),
			},
			expectedValidationErrors: nil,
			expectError:              true,
		},
		{
			name: "No payload provided",
			args: args{
				notificationType: "bonus",
				payload:          nil,
			},
			expectedValidationErrors: nil,
			expectError:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			c := &ValidationClient{
				realTimeSchemas: map[string]gojsonschema.JSONLoader{
					"login_v2": gojsonschema.NewStringLoader(schemaJson),
				},
				operatorAPISchemas:         nil,
				realTimeSchemaRegistry:     nil,
				operatorAPISchemasRegistry: nil,
			}

			actualValidationError, err := c.ValidateRealTimeEvent(tt.args.notificationType, tt.args.payload)
			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, actualValidationError)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedValidationErrors, actualValidationError)
		})
	}
}

func TestClient_ValidateOperatorAPIResponse(t *testing.T) {
	const (
		testValidEvent   = `{"user_id":"123abc","bonus_code":"xyz"}`
		testInvalidEvent = `{"user_id":"123abc","bonus_code":1234,"currency":"USD"}`
		schemaJson       = `{"$schema":"http://json-schema.org/draft-07/schema#","type":"object","description":"SchemafortherequestpayloadofPOST/bonus/credit.Itspecifiestheuseridentifierandabonuscode.","properties":{"user_id":{"type":"string","message":{"format":"Fieldisrequired."},"description":"Theuniqueidentifieroftheuser."},"bonus_code":{"type":"string","message":{"format":"Fieldisrequired."},"description":"Thecoderepresentingaspecificbonus."}},"required":["user_id","bonus_code"],"additionalProperties":false}`
	)

	type args struct {
		endpoint string
		payload  []byte
	}

	tests := []struct {
		name                     string
		args                     args
		expectedValidationErrors []ValidationError
		expectError              bool
	}{
		{
			name: "Valid payload",
			args: args{
				endpoint: "bonus_credit",
				payload:  []byte(testValidEvent),
			},
			expectedValidationErrors: nil,
			expectError:              false,
		},
		{
			name: "Invalid payload",
			args: args{
				endpoint: "bonus_credit",
				payload:  []byte(testInvalidEvent),
			},
			expectedValidationErrors: []ValidationError{
				{
					Path:  "(root)",
					Error: "Additional property currency is not allowed",
				},
				{
					Path:  "(root).bonus_code",
					Error: "Invalid type. Expected: string, given: integer",
				},
			},
			expectError: false,
		},
		{
			name: "Invalid notification type",
			args: args{
				endpoint: "invalid",
				payload:  []byte(testValidEvent),
			},
			expectedValidationErrors: nil,
			expectError:              true,
		},
		{
			name: "No payload provided",
			args: args{
				endpoint: "bonus_credit",
				payload:  nil,
			},
			expectedValidationErrors: nil,
			expectError:              true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ValidationClient{
				realTimeSchemas: nil,
				operatorAPISchemas: map[string]gojsonschema.JSONLoader{
					"bonus_credit": gojsonschema.NewStringLoader(schemaJson),
				},
				realTimeSchemaRegistry:     nil,
				operatorAPISchemasRegistry: nil,
			}

			actualValidationError, err := c.ValidateOperatorAPIResponse(tt.args.endpoint, tt.args.payload)
			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, actualValidationError)
				return
			}

			require.NoError(t, err)
			require.Equal(t, tt.expectedValidationErrors, actualValidationError)

		})
	}
}
