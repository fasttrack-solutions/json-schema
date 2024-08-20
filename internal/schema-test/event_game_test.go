package schemas

import (
	"path/filepath"
	"testing"

	"github.com/fasttrack-solutions/json-schema/pkg/validator"
)

var gameListSchema = filepath.Join(realtTimeSchemaPath, "/game.schema.json")

func loadTestsRealtimePostGame() []schemaTest {
	return []schemaTest{
		{
			name: "Valid data",
			payload: `{
				"game_id": "123-123",
				"name": "Book of dead",
				"slug": "book_of_dead",
				"provider": "netent",
				"category": "table-games",
				"subcategory": "blackjack",
				"supported_devices": ["mobile", "desktop"],
				"last_modified": "2021-01-10T12:00:00Z",
				"is_live": true,
				"origin":["brandname"]
			}`,
			isValid: true,
		},
	}
}

func TestValidData_Realtime_Post_Game(t *testing.T) {
	tests := loadTestsRealtimePostGame()

	runRealtimeTest(t, validator.NotificationTypeGame, tests, nil)
}
