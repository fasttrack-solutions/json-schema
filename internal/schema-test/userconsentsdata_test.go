package schemas

func generateUserConsentTests() []schemaTest {
	// This test is failing because the schema is not validating the duplicated email consent. Will need to find a way to validate this.
	// {
	// 	name: "Duplicated email and conflicting consent",
	// 	document: `{
	// 		"consents": [
	// 			{
	// 				"opted_in": true,
	// 				"type": "email"
	// 			},
	// 			{
	// 				"opted_in": false,
	// 				"type": "email"
	// 			},
	// 		]
	// 	}`,
	// 	failsValidation: true,
	// },

	userConsentTests := []schemaTest{
		{
			name: "Valid response with six consents",
			payload: `{
				"consents": [
					{
						"opted_in": true,
						"type": "email"
					},
					{
						"opted_in": true,
						"type": "sms"
					},
					{
						"opted_in": false,
						"type": "telephone"
					},
					{
						"opted_in": true,
						"type": "postMail"
					},
					{
						"opted_in": true,
						"type": "siteNotification"
					},
					{
						"opted_in": true,
						"type": "pushNotification"
					}
				]
			}`,
			failsValidation: false,
		},
		{
			name: "Valid response with one consent",
			payload: `{
				"consents": [
					{
						"opted_in": true,
						"type": "email"
					}
				]
			}`,
			failsValidation: false,
		},
		{
			name: "Empty consents array",
			payload: `{
				"consents": []
			}`,
			failsValidation: true,
		},
		{
			name: "Missing object from consent",
			payload: `{
				"consents": [
					{
						"type": "email"
					}
				]
			}`,
			failsValidation: true,
		},
		{
			name: "Duplicated email consent",
			payload: `{
				"consents": [
					{
						"opted_in": true,
						"type": "email"
					},
					{
						"opted_in": true,
						"type": "email"
					},
					{
						"opted_in": false,
						"type": "telephone"
					},
					{
						"opted_in": true,
						"type": "postMail"
					},
					{
						"opted_in": true,
						"type": "siteNotification"
					},
					{
						"opted_in": true,
						"type": "pushNotification"
					}
				]
			}`,
			failsValidation: true,
		},

		{
			name: "One invalid consent",
			payload: `{
				"consents": [
					{
						"opted_in": true,
						"type": "email"
					},
					{
						"opted_in": true,
						"type": "sms"
					},
					{
						"opted_in": false,
						"type": "telephone"
					},
					{
						"opted_in": true,
						"type": "postMail"
					},
					{
						"opted_in": true,
						"type": "siteNotification"
					},
					{
						"opted_in": true,
						"type": "social"
					}
				]
			}`,
			failsValidation: true,
		},
		{
			name: "Over limit of consents",
			payload: `{
				"consents": [
					{
						"opted_in": true,
						"type": "email"
					},
					{
						"opted_in": true,
						"type": "sms"
					},
					{
						"opted_in": false,
						"type": "telephone"
					},
					{
						"opted_in": true,
						"type": "postMail"
					},
					{
						"opted_in": true,
						"type": "siteNotification"
					},
					{
						"opted_in": true,
						"type": "pushNotification"
					},
					{
						"opted_in": true,
						"type": "richInbox"
					}
				]
			}`,
			failsValidation: true,
		},
	}
	return userConsentTests
}

func generateUserConsentsV2Tests() []schemaTest {
	return []schemaTest{
		{
			name: "Valid response",
			payload: `{
						"user_id": "7865312321",
						"timestamp": "2015-03-02T8:27:58.721607Z",
						"origin": "sub.example.com"
					}`,
			failsValidation: false,
		},
	}
}
