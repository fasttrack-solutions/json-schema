package schemas

import (
	"path/filepath"
	"testing"
)

var casinoSchema = filepath.Join(realtTimeSchemaPath, "/casino.schema.json")

func loadTestsRealtimeCasinoBet() []schemaTest {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "2019020103308257480",
				"amount": 20,
				"balance_after": 174.03,
				"balance_before": 194.03,
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"game_id": "10911",
				"game_name": "Dead or Alive",
				"game_type": "Slots",
				"is_round_end": false,
				"locked_wager_amount": 0.00,
				"origin": "sub.example.com",
				"round_id": "2019020103308257480",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607Z",
				"type": "Bet",
				"user_id": "52530",
				"vendor_id": 123,
				"vendor_name": "Netent",
				"wager_amount": 20,
				"meta":{
					"key1": 10.00,
					"key2": "some string",
					"key3": false
				}
			}`,
			isValid: true,
		},
	}
	return tests
}

func loadTestsRealtimeCasinoWin() []schemaTest {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"activity_id": "2019020103308257481",
				"amount": 5,
				"balance_after": 195.83,
				"balance_before": 190.83,
				"bonus_wager_amount": 0.0,
				"currency": "EUR",
				"exchange_rate": 1,
				"game_id": "10911",
				"game_name": "Dead or Alive",
				"game_type": "Slots",
				"is_round_end": true,
				"locked_wager_amount": 0.00,
				"origin": "sub.example.com",
				"round_id": "12345",
				"status": "Approved",
				"timestamp": "2015-03-02T8:27:58.721607Z",
				"type": "Win",
				"user_id": "52530",
				"vendor_id": 123,
				"vendor_name": "Netent",
				"wager_amount": 5,
				"meta":{
					"key1": 10.00,
					"key2": "some string",
					"key3": false
				}
			}`,
			isValid: true,
		},
	}
	return tests
}

func TestValidData_RealtimeCasinoBet(t *testing.T) {
	runTestsRealtimeCasino(t, casinoSchema, loadTestsRealtimeCasinoBet(), EventEnums{
		typeEnums: []string{"Bet"},
	})
}

func TestValidData_RealtimeCasinoWin(t *testing.T) {
	runTestsRealtimeCasino(t, casinoSchema, loadTestsRealtimeCasinoWin(), EventEnums{
		typeEnums: []string{"Bet", "Win"},
	})
}

func runTestsRealtimeCasino(t *testing.T, schema string, schemaTests []schemaTest, enums EventEnums) {
	runTestCases(t, schema, schemaTests, &enums)
}
