package validator

import (
	"errors"
	"fmt"
)

const pathSchemasRealTime = "schemas/real-time-events"

var notificationTypes = []string{
	"bonus",
	"cart",
	"casino",
	"custom",
	"game",
	"login_v2",
	"lottery_v2",
	"lottery",
	"payment",
	"sportsbook",
	"user_balances_update",
	"user_block_v2",
	"user_consents_v2",
	"user_create_v2",
	"user_update_v2",
}

// ValidateRealTimeEvent takes in notificationType which can be the notification type for the real time feed event, and payload which is the json payload.
// Allowed notification types are: bonus, casino, custom, game, login_v2, lottery_v2, payment, sportsbook, user_balances_update, user_block_v2, user_consents_v2, user_create_v2, user_update_v2
func (c *Client) ValidateRealTimeEvent(notificationType string, payload []byte) ([]Error, error) {
	if notificationType == "" {
		return nil, errors.New("no payload type provided")
	}

	if !isValidType(notificationType, notificationTypes) {
		return nil, fmt.Errorf("%s is an invalid notification type", notificationType)
	}

	errors, err := validatePayload(notificationType, payload, c.eventSchemas)
	if err != nil {
		return nil, fmt.Errorf("validating %s event json payload: %v", notificationType, err)
	}

	return errors, nil
}
