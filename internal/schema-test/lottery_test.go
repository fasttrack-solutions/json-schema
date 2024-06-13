package schemas

import (
	"path/filepath"
	"testing"
)

var lotterySchema = filepath.Join(realtTimeSchemaPath, "/lottery.schema.json")
var lotteryV2Schema = filepath.Join(realtTimeSchemaPath, "/lottery_v2.schema.json")
var cartSchema = filepath.Join(realtTimeSchemaPath, "/cart.schema.json")

func TestValidData_Realtime_Post_LotteryPurchase(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"ticket_id": "2019020103308257480",
				"lottery_id": "1",
				"lottery_name": "Mega Millions",  
				"lottery_type": "Lottery",
				"draw_id": "1",
				"draw_date": "2015-03-02T8:27:58.721607+06:00",
				"currency": "EUR",
				"exchange_rate": 1,
				"bonus_code": "Welcome",
				"amount": 4.0,
				"discount_amount": 2.0,
				"lines": [
					{
						"line_id": 1,
						"is_free_bet": false,
						"numbers": [12,31,32,11,35,54],
						"special_numbers": [35],
						"is_quick_pick": true,
						"meta": {}
					},
					{
						"line_id": 2,
						"is_free_bet": true,
						"numbers": [32,23,12,11,35,54],
						"special_numbers": [32],
						"is_quick_pick": false,
						"meta": {}
					}
				],
				"origin": "sub.example.com",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Bet",
				"user_id": "52530",
				"device_type": "Mobile",
				"meta": {
					"key1": 10.00,
					"key2": "some string",
					"key3": false
				}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, lotteryV2Schema, tests, &EventEnums{
		typeEnums: []string{"Bet"},
	})
}

func TestValidData_Realtime_Post_LotterySettlement(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"ticket_id": "2019020103308257480",
				"lottery_id": "1",
				"lottery_name": "Mega Millions",  
				"lottery_type": "Lottery",
				"draw_id": "1",
				"draw_date": "2015-03-02T8:27:58.721607+06:00",
				"currency": "EUR",
				"exchange_rate": 1,
				"draw_numbers": [32,23,12,11,35,54],
				"draw_special_numbers": [32],
				"amount": 2400,
				"lines": [
					{
						"line_id": 1,
						"matched_numbers": [12,32,35,54],
						"is_winning_line": false,
						"meta": {}
					},
					{
						"line_id": 2,
						"matched_numbers": [32,23,12,11,35,54],
						"matched_special_numbers": [32],
						"is_winning_line": true,
						"meta": {}
					}
				],
				"origin": "sub.example.com",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Settlement",
				"user_id": "52530",
				"meta": {}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, lotteryV2Schema, tests, &EventEnums{
		typeEnums: []string{"Settlement"},
	})
}

func TestValidData_Realtime_Post_LotterySubscription(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"timestamp": "2019-02-18T10:35:03-08:00",
				"transaction_id": "23767444",
				"user_id": "5288704",
				"amount": 36,
				"account_amount": 16.80,
				"bonus_amount": 0,
				"bonus_code": "",
				"currency": "BRL",
				"deposit_amount": 19.20,
				"exchange_rate": 1,
				"origin": "sub.example.com",
				"status": "Approved",
				"tickets": [
					{
					"ticket_id": "17167745",
					"product": "SingleTicket",
					"product_name": "Mega-Sena",
					"amount": 18,
					"lines_count": 1,
					"draws": [{ "timestamp": "2019-02-20T14:00:00-07:00" }]
					},
					{
					"ticket_id": "17167746",
					"product": "SingleTicket",
					"product_name": "Mega-Sena",
					"amount": 18,
					"lines_count": 1,
					"draws": [{ "timestamp": "2019-02-20T14:00:00-07:00" }]
					}
				]
			}`,
			isValid: true,
		},
	}

	runTestCases(t, lotterySchema, tests, &EventEnums{})
}

func TestValidData_Realtime_Post_LotteryCart(t *testing.T) {
	tests := []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"amount": 100.0,
				"bonus_code": "",
				"cart_id": 1,
				"discount_amount": 10.00,
				"status": "Successful",
				"timestamp": "2015-03-02T8:27:58.721607+06:00",
				"type": "Checkout",
				"user_id": "abc1234",
				"currency": "EUR",
				"device_type": "mobile",
				"exchange_rate": 1.00,
				"origin": "sub.example.com" , 
				"payments": [
					{
						"type": "Wallet",
						"amount": 70.00,
						"meta": {}
					},
					{
						"type": "Deposit",
						"amount": 20.00,
						"meta": {}
					},
					{
						"type": "VirtualCurrency",
						"amount": 0.00,
						"meta": {}
					}
				],
				"items": [
					{
						"id": 1,
						"type": "Scratch",
						"name": "Scratch And Win",
						"description": "Win this summers biggest prize",
						"amount": 10.00,
						"discount_amount": 0.00,
						"bonus_code": "",
						"meta": {}
					},
					{
						"id": 2,
						"type": "Lottery",
						"name": "Mega Millions",
						"description": "Weekend Lottery Lines",
						"amount": 90.00,
						"discount_amount": 0.00,
						"bonus_code": "",
						"meta": {}
					}
				],
				"meta":{
					"key1": 10.00,
					"key2": "some string",
					"key3": false
				}
			}`,
			isValid: true,
		},
	}

	runTestCases(t, cartSchema, tests, &EventEnums{
		statusEnums: []string{"Successful"},
		typeEnums:   []string{"Checkout"},
	})
}
