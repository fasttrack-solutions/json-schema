package schemas

import (
	"path/filepath"
	"testing"
)

var pathSportsbook = filepath.Join(realtTimeSchemaPath, "/sportsbook.schema.json")

func TestValidData_Realtime_Get_SportsSingle(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "101212",
				"activity_id_reference": "",
				"amount": 100,
				"balance_after": 852.21,
				"balance_before": 952.21,
				"bet_type": "Single",
				"odds": 6.25,
				"bets": [
					{
					"event_name": "Queens Park Rangers vs West Bromwich Albion",
					"is_free_bet": false,
					"is_risk_free_bet": false,
					"market": "Correct Score",
					"match_start": "2015-03-02T8:27:58.721607+06:00",
					"outcomes": [
						{
						"criterion_name": "Correct Score",
						"outcome_label": "1-0",
						"meta": {}
						}
					],
					"sports_name": "Football",
					"tournament_name": "Championship",
					"odds": 6.25,
					"is_live": true,
					"meta": {}
					}
				],
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"is_cashout": false,
				"locked_wager_amount": 0.0,
				"origin": "sub.example.com",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Bet",
				"user_id": "12304",
				"wager_amount": 0.0,
				"device_type": "App",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, pathSportsbook, tests, nil)
}

func TestValidData_Realtime_Get_SportsMulti(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "101213",
				"activity_id_reference": "",
				"amount": 50,
				"balance_after": 950,
				"balance_before": 1000,
				"bet_type": "Multi",
				"odds": 10.19,
				"bets": [
					{
					"event_name": "Dortmund vs Tottenham",
					"is_free_bet": false,
					"is_risk_free_bet": false,
					"market": "Correct Score",
					"match_start": "2015-03-02T8:27:58.721607+06:00",
					"outcomes": [
						{
							"criterion_name": "Correct Score",
							"outcome_label": "1-0",
							"meta": {}
						}
					],
					"sports_name": "Football",
					"tournament_name": "Champions League",
					"odds": 1.95,
					"is_live": false,
					"meta": {}
					},
					{
					"event_name": "FC Schalke 04 vs Manchester City",
					"is_free_bet": false,
					"is_risk_free_bet": false,
					"market": "1X2",
					"match_start": "2015-03-02T8:27:58.721607+06:00",
					"outcomes": [
						{
							"criterion_name": "1X2",
							"outcome_label": "1",
							"meta": {}
						}
					],
					"sports_name": "Football",
					"tournament_name": "Champions League",
					"odds": 1.26,
					"is_live": false,
					"meta": {}
					},
					{
					"event_name": "Olympique Lyon vs FC Barcelona",
					"is_free_bet": false,
					"is_risk_free_bet": false,
					"market": "1X2",
					"match_start": "2015-03-02T8:27:58.721607+06:00",
					"outcomes": [
						{
							"criterion_name": "1X2",
							"outcome_label": "2",
							"meta": {}
						}
					],
					"sports_name": "Football",
					"tournament_name": "Champions League",
					"odds": 4.15,
					"is_live": false,
					"meta": {}
					}
				],
				"bonus_wager_amount": 0,
				"currency": "EUR",
				"exchange_rate": 1,
				"is_cashout": true,
				"locked_wager_amount": 3.13,
				"origin": "sub.example.com",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Bet",
				"user_id": "12304",
				"wager_amount": 0.0,
				"device_type": "App",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, pathSportsbook, tests, nil)
}

func TestValidData_Realtime_Get_SportsSettlment(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "101420",
				"activity_id_reference": "101212",
				"amount": 625.0,
				"balance_after": 1477.21,
				"balance_before": 852.21,
				"bet_type": "Single",
				"bets": [],
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"is_cashout": false,
				"locked_wager_amount": 0.0,
				"origin": "sub.example.com",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Settlement",
				"user_id": "12304",
				"wager_amount": 0.0,
				"device_type": "App",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, pathSportsbook, tests, nil)
}

func TestValidData_Realtime_Get_SportsCashout(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "101422",
				"activity_id_reference": "101212",
				"amount": 262.20,
				"balance_after": 1214.41,
				"balance_before": 952.21,
				"bet_type": "Single",
				"bets": [],
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"is_cashout": true,
				"locked_wager_amount": 0.0,
				"origin": "sub.example.com",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Settlement",
				"user_id": "12304",
				"wager_amount": 0.0,
				"device_type": "App",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, pathSportsbook, tests, nil)
}

func TestValidData_Realtime_Get_SportsCashoutMulti(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "101529",
				"activity_id_reference": "101213",
				"amount": 26.40,
				"balance_after": 976.4,
				"balance_before": 950,
				"bet_type": "Multi",
				"bets": [],
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"is_cashout": true,
				"locked_wager_amount": 0.0,
				"origin": "sub.example.com",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Settlement",
				"user_id": "12304",
				"wager_amount": 0.0,
				"device_type": "App",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, pathSportsbook, tests, nil)
}
