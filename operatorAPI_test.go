package json_schema

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var client *APIClient

func TestMain(m *testing.M) {
	authToken := os.Getenv("AUTH_TOKEN")

	testServer := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(handlePayload(r.URL.Path))
		}),
	)

	client = NewAPIClient(testServer.URL, authToken)

	m.Run()
}

func TestUserDetails(t *testing.T) {
	client.TestEndpoint(t, "/userdetails/59924", http.StatusOK, "operator-api/get/user_details.schema.json")
}

func TestUserBlocks(t *testing.T) {
	client.TestEndpoint(t, "/userblocks/59924", http.StatusOK, "operator-api/get/user_blocks.schema.json")
}

func TestGetUserConsents(t *testing.T) {
	client.TestEndpoint(t, "/userconsents/59924", http.StatusOK, "operator-api/get/user_consents.schema.json")
}

func handlePayload(path string) json.RawMessage {
	switch path {
	case "/userdetails/59924":
		return json.RawMessage(`{"address":"Tower Road, 120A","affiliate_reference":"AFF_1234A_UK","birth_date":"2015-03-02","city":"Sliema","country":"MT","currency":"EUR","email":"tony@example.com","first_name":"Tony","is_blocked":false,"is_excluded":false,"language":"en","last_name":"Carrot","market":"gb","mobile":"21435678","mobile_prefix":"+356","origin":"sub.example.com","postal_code":"SLM 1030","registration_code":"ABC123","registration_date":"2015-03-02T08:27:58.721607Z","roles":["VIP","TEST_USER"],"sex":"male","title":"Mr","user_id":"7865312321","username":"PirateTony34","verified_at":"2015-03-02T08:27:58.721607Z","segmentation":{"vip_level":15,"special_segmentation":"3D"}}`)

	case "/userblocks/59924":
		return json.RawMessage(`{"blocks":[{"active":true,"type":"Excluded","note":"Exclusion reason"},{"active":true,"type":"Blocked","note":"Block reason"}]}`)

	case "/userconsents/59924":
		return json.RawMessage(`{"consents":[{"opted_in":true,"type":"email"},{"opted_in":true,"type":"sms"},{"opted_in":false,"type":"telephone"},{"opted_in":true,"type":"postMail"},{"opted_in":true,"type":"siteNotification"},{"opted_in":true,"type":"pushNotification"}]}`)
	}

	return nil
}
