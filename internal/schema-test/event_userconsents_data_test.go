package schemas

func generateUserConsentTests() []schemaTest {
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
			isValid: true,
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
			isValid: true,
		},
		{
			name: "Empty consents array",
			payload: `{
				"consents": []
			}`,
			isValid: false,
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
			isValid: false,
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
			isValid: false,
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
			isValid: false,
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
			isValid: false,
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
			isValid: true,
		},
	}
}
