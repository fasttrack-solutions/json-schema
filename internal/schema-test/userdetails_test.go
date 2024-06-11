package schemas

import (
	"path/filepath"
	"testing"
)

var pathUserDetails = filepath.Join(operatorApiSchemaPath, "/get/user_details.schema.json")

func TestValidData_Operator_Get_UserDetails(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"address": "Ratatouille avenue, 34B",
				"birth_date": "2015-03-02",
				"city": "Paris",
				"country": "MT",
				"currency": "USD",
				"email": "tony@example.com",
				"first_name": "Tony",
				"is_blocked": false,
				"is_excluded": false,
				"language": "en",
				"last_name": "Carrot",
				"mobile": "2143567",
				"mobile_prefix": "256",
				"origin": "sub.example.com",
				"postal_code": "TTS52",
				"roles": [
					"VIP",
					"TEST_USER"
				],
				"sex": "female",
				"title": "Dr",
				"user_id": "7865312321",
				"username": "PirateTony34",
				"verified_at": "2015-03-02T08:27:58.721607Z",
				"registration_code": "ABC123",
				"registration_date": "2015-03-02T08:27:58.721607Z",
				"affiliate_reference": "AFF_1234A_UK",
				"market": "en",
				"segmentation": {
					"vip_level": 15,
					"special_segmentation": "3D"
				}
			}`,
			failsValidation: false,
		},
	}

	runTests(t, pathUserDetails, tests)
}
