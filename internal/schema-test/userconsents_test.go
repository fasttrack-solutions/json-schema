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
	runTests(t, pathUserConsents, generateUserConsentTests())
}

func TestValidData_Operator_Post_UserConsents(t *testing.T) {
	runTests(t, pathPostUserConsents, generateUserConsentTests())
}

func TestValidData_Realtime_Put_UserConsents(t *testing.T) {
	userConsentsV2Tests := generateUserConsentsV2Tests()

	runTests(t, pathUserConsentsV2, userConsentsV2Tests)
	validateUserIDField(t, pathUserConsentsV2, userConsentsV2Tests[0])
	validateTimestampField(t, pathUserConsentsV2, userConsentsV2Tests[0])
	validateOriginField(t, pathUserConsentsV2, userConsentsV2Tests[0])
}
