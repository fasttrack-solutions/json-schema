package json_schema

import (
	"net/http"
	"os"
	"testing"
)

var client *APIClient

func TestMain(m *testing.M) {
	baseURL := os.Getenv("BASE_URL")
	authToken := os.Getenv("AUTH_TOKEN")
	client = NewAPIClient(baseURL, authToken)

	m.Run()
}

func TestUserDetails(t *testing.T) {
	client.TestEndpoint(t, "/userdetails/59924", http.StatusOK, "user_details_json_schema.json")
}

func TestUserBlocks(t *testing.T) {
	client.TestEndpoint(t, "/userblocks/59924", http.StatusOK, "user_blocks_json_schema.json")
}

func TestGetUserConsents(t *testing.T) {
	client.TestEndpoint(t, "/userconsents/59924", http.StatusOK, "user_consents_GET_json_schema.json")
}
