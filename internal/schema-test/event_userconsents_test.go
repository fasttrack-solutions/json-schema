package schemas

import (
	"path/filepath"
	"testing"
)

var (
	pathUserConsents     = filepath.Join(operatorApiSchemaPath, "/get/user_consents.schema.json")
	pathPostUserConsents = filepath.Join(operatorApiSchemaPath, "/post/post_user_consents.schema.json")
	pathUserConsentsV2   = filepath.Join(realtTimeSchemaPath, "/user_consents_v2.schema.json")
)

func TestValidData_Operator_Get_UserConsents(t *testing.T) {
	runOperatorTest(t, pathUserConsents, generateUserConsentTests(), nil)
}

func TestValidData_Operator_Post_UserConsents(t *testing.T) {
	runOperatorTest(t, pathPostUserConsents, generateUserConsentTests(), nil)
}

func TestValidData_Realtime_Put_UserConsents(t *testing.T) {
	runRealtimeTest(t, pathUserConsentsV2, generateUserConsentsV2Tests(), nil)
}
