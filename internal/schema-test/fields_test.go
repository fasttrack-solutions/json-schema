package schemas

import (
	"testing"
)

const keyCurrency = "currency"
const keyGameID = "game_id"
const keyOrigin = "origin"
const keyStatus = "status"
const keyUserID = "user_id"

func validateCurrencyField(t *testing.T, schemaPath string, test schemaTest) []schemaTest {
	return validateField(t, "Currency", keyCurrency, schemaPath, test, loadFieldTestsString("EUR", false))
}

func validateOriginField(t *testing.T, schemaPath string, test schemaTest) []schemaTest {
	return validateField(t, "Origin", keyOrigin, schemaPath, test, loadFieldTestsString("Origin", false))
}

func validateStatusField(t *testing.T, schemaPath string, test schemaTest) []schemaTest {
	return validateField(t, "Status", keyStatus, schemaPath, test, loadFieldTestsString("Pending", false))
}

func validateUserIDField(t *testing.T, schemaPath string, test schemaTest) []schemaTest {
	return validateField(t, "User ID", keyUserID, schemaPath, test, loadFieldTestsString("UserID", false))
}

func validateGameIDField(t *testing.T, schemaPath string, test schemaTest) []schemaTest {
	return validateField(t, "Game ID", keyGameID, schemaPath, test, loadFieldTestsString("Origin", false))
}
