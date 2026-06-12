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
			name: "Valid payload with uppercase notification type",
			args: args{
				notificationType: "LOGIN_V2",
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

			c := &Client{
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

func TestClient_ValidateRealTimeEvent_ITPing(t *testing.T) {
	// Use real embedded schemas via NewClient to verify the it_ping schema loads and validates correctly.
	c, err := NewClient()
	require.NoError(t, err)

	tests := []struct {
		name                     string
		notificationType         string
		payload                  string
		expectedValidationErrors []ValidationError
		expectError              bool
	}{
		{
			name:             "Valid IT_PING payload",
			notificationType: "IT_PING",
			payload:          `{"ping_id":"6e16158b-df22-4a52-9a5d-12ff1b71da09","user_id":"-1","timestamp":"2026-02-03T19:44:35Z","origin":"defaultorigin"}`,
		},
		{
			name:             "Valid it_ping payload (lowercase type)",
			notificationType: "it_ping",
			payload:          `{"ping_id":"abc-123","user_id":"42","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			name:             "Missing required field ping_id",
			notificationType: "IT_PING",
			payload:          `{"user_id":"-1","timestamp":"2026-02-03T19:44:35Z","origin":"defaultorigin"}`,
			expectedValidationErrors: []ValidationError{
				{
					Path:  "(root)",
					Error: "ping_id is required",
				},
			},
		},
		{
			name:             "Missing multiple required fields",
			notificationType: "IT_PING",
			payload:          `{"ping_id":"abc-123"}`,
			expectedValidationErrors: []ValidationError{
				{
					Path:  "(root)",
					Error: "origin is required",
				},
				{
					Path:  "(root)",
					Error: "timestamp is required",
				},
				{
					Path:  "(root)",
					Error: "user_id is required",
				},
			},
		},
		{
			name:             "Invalid user_id type (integer instead of string)",
			notificationType: "IT_PING",
			payload:          `{"ping_id":"abc-123","user_id":42,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{
					Path:  "(root).user_id",
					Error: "Invalid type. Expected: string, given: integer",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationErrors, err := c.ValidateRealTimeEvent(tt.notificationType, []byte(tt.payload))
			if tt.expectError {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			require.Equal(t, tt.expectedValidationErrors, validationErrors)
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
			name: "Valid payload with uppercase endpoint",
			args: args{
				endpoint: "BONUS_CREDIT",
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
			c := &Client{
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

// TestClient_ValidateRealTimeEvent_NewEventSchemas verifies that every event
// schema added for DEV-17460 loads via NewClient and accepts a valid payload.
func TestClient_ValidateRealTimeEvent_NewEventSchemas(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)

	tests := []struct {
		notificationType string
		payload          string
	}{
		{
			notificationType: "USER_CREATE",
			payload:          `{"user_id":"123","username":"PirateTony34","email":"example@site.com","currency":"USD","roles":["VIP"],"allows_email_marketing":true,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "USER_UPDATE",
			payload:          `{"user_id":"123","first_name":"Tony","is_blocked":false,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "USER_CONSENTS",
			payload:          `{"user_id":"123","consents":[{"type":"Email","opted_in":true,"products":["casino"]}],"action":"site-update","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "LOGIN",
			payload:          `{"user_id":"123","is_impersonated":false,"ip_address":"192.0.2.1","user_agent":"Mozilla/5.0","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "GAME_ROUND",
			payload:          `{"user_id":"123","round_id":"r1","game_id":"g1","game_name":"Starburst","game_type":"slots","vendor_id":"v1","user_currency":"EUR","device_type":"mobile","real_bet_user":1.5,"real_win_user":0,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "USER_BLOCK",
			payload:          `{"user_id":"123","type":"self-exclusion","note":"requested by user","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "GLOBAL_EVENT",
			payload:          `{"origin":"test","notification_type":"jackpot-won","timestamp":"2026-01-01T00:00:00Z","metadata":{"jackpot_id":"j1"}}`,
		},
		{
			notificationType: "MARKETING_VERIFIED",
			payload:          `{"user_id":"123","verified":true,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			// verified is a pointer in the Go model and may be explicitly null.
			notificationType: "MARKETING_VERIFIED",
			payload:          `{"user_id":"123","verified":null,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "CASUMO_VALUABLE",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","slug":"golden-chest","title":"Golden Chest","description":"A shiny chest","status":"Fresh"}`,
		},
		{
			notificationType: "CASUMO_BELT",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","id":"b1","name":"Blue Belt"}`,
		},
		{
			notificationType: "CASUMO_LEVEL_UP",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","new_level":3,"previous_level":2}`,
		},
		{
			notificationType: "USER_DEVICE_LINK",
			payload:          `{"user_id":"123","tokens":[{"token":"abc","channel":"ios","provider":"firebase"}],"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "user_device_link",
			payload:          `{"user_id":"123","tokens":[],"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
		},
		{
			notificationType: "USER_DEVICE_UNLINK",
			payload:          `{"user_id":"123","tokens":[{"token":"abc","channel":"web","provider":"firebase"}],"timestamp":"2026-01-01T00:00:00Z","origin":"test","skip_publishing":true}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.notificationType, func(t *testing.T) {
			validationErrors, err := c.ValidateRealTimeEvent(tt.notificationType, []byte(tt.payload))
			require.NoError(t, err)
			require.Empty(t, validationErrors)
		})
	}
}

// TestClient_ValidateRealTimeEvent_NewEventSchemas_Invalid covers one
// representative validation failure per event schema added for DEV-17460.
func TestClient_ValidateRealTimeEvent_NewEventSchemas_Invalid(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)

	tests := []struct {
		name                     string
		notificationType         string
		payload                  string
		expectedValidationErrors []ValidationError
	}{
		{
			name:             "USER_CREATE with integer user_id",
			notificationType: "USER_CREATE",
			payload:          `{"user_id":123,"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root).user_id", Error: "Invalid type. Expected: string, given: integer"},
			},
		},
		{
			name:             "USER_UPDATE missing origin",
			notificationType: "USER_UPDATE",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "origin is required"},
			},
		},
		{
			name:             "USER_CONSENTS missing consents",
			notificationType: "USER_CONSENTS",
			payload:          `{"user_id":"123","action":"site-update","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "consents is required"},
			},
		},
		{
			name:             "LOGIN missing timestamp",
			notificationType: "LOGIN",
			payload:          `{"user_id":"123","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "timestamp is required"},
			},
		},
		{
			name:             "GAME_ROUND missing device_type",
			notificationType: "GAME_ROUND",
			payload:          `{"user_id":"123","round_id":"r1","game_id":"g1","game_name":"Starburst","game_type":"slots","vendor_id":"v1","user_currency":"EUR","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "device_type is required"},
			},
		},
		{
			name:             "USER_BLOCK missing type",
			notificationType: "USER_BLOCK",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "type is required"},
			},
		},
		{
			name:             "GLOBAL_EVENT missing notification_type",
			notificationType: "GLOBAL_EVENT",
			payload:          `{"origin":"test","timestamp":"2026-01-01T00:00:00Z"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "notification_type is required"},
			},
		},
		{
			name:             "MARKETING_VERIFIED missing verified",
			notificationType: "MARKETING_VERIFIED",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "verified is required"},
			},
		},
		{
			name:             "CASUMO_VALUABLE missing status",
			notificationType: "CASUMO_VALUABLE",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","slug":"golden-chest","title":"Golden Chest","description":"A shiny chest"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "status is required"},
			},
		},
		{
			name:             "CASUMO_BELT missing id",
			notificationType: "CASUMO_BELT",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","name":"Blue Belt"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "id is required"},
			},
		},
		{
			name:             "CASUMO_LEVEL_UP with string new_level",
			notificationType: "CASUMO_LEVEL_UP",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test","new_level":"3"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root).new_level", Error: "Invalid type. Expected: integer, given: string"},
			},
		},
		{
			name:             "USER_DEVICE_LINK missing tokens",
			notificationType: "USER_DEVICE_LINK",
			payload:          `{"user_id":"123","timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root)", Error: "tokens is required"},
			},
		},
		{
			name:             "USER_DEVICE_LINK with invalid channel",
			notificationType: "USER_DEVICE_LINK",
			payload:          `{"user_id":"123","tokens":[{"token":"abc","channel":"windows","provider":"firebase"}],"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root).tokens.0.channel", Error: `tokens.0.channel must be one of the following: "android", "ios", "web", "unknown"`},
			},
		},
		{
			name:             "USER_DEVICE_UNLINK with missing token provider",
			notificationType: "USER_DEVICE_UNLINK",
			payload:          `{"user_id":"123","tokens":[{"token":"abc","channel":"ios"}],"timestamp":"2026-01-01T00:00:00Z","origin":"test"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root).tokens.0", Error: "provider is required"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			validationErrors, err := c.ValidateRealTimeEvent(tt.notificationType, []byte(tt.payload))
			require.NoError(t, err)
			require.Equal(t, tt.expectedValidationErrors, validationErrors)
		})
	}
}

// TestClient_Validate_SchemaNotFound verifies that unknown payload types are
// reported via the ErrSchemaNotFound sentinel so callers can fall back
// gracefully instead of treating the payload as unprocessable (DEV-17460).
func TestClient_Validate_SchemaNotFound(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)

	t.Run("unknown real time event type", func(t *testing.T) {
		validationErrors, err := c.ValidateRealTimeEvent("BOGUS_TYPE", []byte(`{}`))
		require.ErrorIs(t, err, ErrSchemaNotFound)
		require.ErrorContains(t, err, "invalid payload type bogus_type")
		require.Nil(t, validationErrors)
	})

	t.Run("empty real time event type", func(t *testing.T) {
		validationErrors, err := c.ValidateRealTimeEvent("", []byte(`{}`))
		require.ErrorIs(t, err, ErrSchemaNotFound)
		require.Nil(t, validationErrors)
	})

	t.Run("unknown operator API endpoint", func(t *testing.T) {
		validationErrors, err := c.ValidateOperatorAPIResponse("unknown_endpoint", []byte(`{}`))
		require.ErrorIs(t, err, ErrSchemaNotFound)
		require.Nil(t, validationErrors)
	})

	t.Run("nil payload for known type is not ErrSchemaNotFound", func(t *testing.T) {
		validationErrors, err := c.ValidateRealTimeEvent("IT_PING", nil)
		require.Error(t, err)
		require.NotErrorIs(t, err, ErrSchemaNotFound)
		require.Nil(t, validationErrors)
	})
}
