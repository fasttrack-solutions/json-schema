package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/xeipuuv/gojsonschema"
)

// These schemas must accept what the services actually send: crm-integration-api's GameRequest
// (Origins []string `json:"origin"`), its SportsRequest (odds has no binding:"required"), and
// sms-distributor's operator SMS payload (`json:"encoding"`).
func TestClient_ValidateRealTimeEvent_MatchesServiceModels(t *testing.T) {
	c, err := NewClient()
	require.NoError(t, err)

	tests := []struct {
		name                     string
		notificationType         string
		payload                  string
		expectedValidationErrors []ValidationError
	}{
		{
			name:             "Game sync with origin as a list",
			notificationType: "GAME",
			payload:          `{"game_id":"7elements","name":"7 Elements","is_live":true,"origin":["SLOTEXPRESS"],"supported_devices":["desktop","mobile"],"rtp":96}`,
		},
		{
			name:             "Game sync with origin as a string is rejected",
			notificationType: "GAME",
			payload:          `{"game_id":"7elements","name":"7 Elements","is_live":true,"origin":"SLOTEXPRESS"}`,
			expectedValidationErrors: []ValidationError{
				{Path: "(root).origin", Error: "Invalid type. Expected: array, given: string"},
			},
		},
		{
			name:             "Sportsbook settlement without odds",
			notificationType: "SPORTSBOOK",
			payload:          `{"activity_id":"101420","activity_id_reference":"101212","amount":625.0,"bet_type":"Single","bets":[],"currency":"EUR","exchange_rate":1,"is_cashout":false,"origin":"test","status":"Approved","timestamp":"2026-03-02T08:27:58Z","type":"Settlement","user_id":"1"}`,
		},
		{
			name:             "Sportsbook cashout without odds",
			notificationType: "SPORTSBOOK",
			payload:          `{"activity_id":"101422","activity_id_reference":"101212","amount":262.2,"bet_type":"Single","bets":[],"currency":"EUR","exchange_rate":1,"is_cashout":true,"origin":"test","status":"Approved","timestamp":"2026-03-02T08:27:58Z","type":"Settlement","user_id":"1"}`,
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

func TestCustomSMSSchema_UsesEncoding(t *testing.T) {
	raw, err := schemasFileSystem.ReadFile("schemas/operator-api/post/custom_sms.schema.json")
	require.NoError(t, err)
	schema, err := gojsonschema.NewSchema(gojsonschema.NewBytesLoader(raw))
	require.NoError(t, err)

	const activity = `"activity":{"id":"1","brand_id":1,"user_id":"1","internal_user_id":1,"activity_id":1,"action_group_id":1,"trigger_hash":"h","action_id":1}`

	result, err := schema.Validate(gojsonschema.NewStringLoader(`{` + activity + `,"content":"hi","recipient":"+35699000000","sender_name":"FT","provider_name":"operator_api","encoding":"auto"}`))
	require.NoError(t, err)
	require.True(t, result.Valid(), "%v", result.Errors())

	result, err = schema.Validate(gojsonschema.NewStringLoader(`{` + activity + `,"content":"hi","recipient":"+35699000000","sender_name":"FT","provider_name":"operator_api","enconding":"auto"}`))
	require.NoError(t, err)
	require.False(t, result.Valid())
}
