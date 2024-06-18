package json_schema

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/xeipuuv/gojsonschema"
)

type APIClient struct {
	baseURL   string
	authToken string
}

func NewAPIClient(baseURL string, authToken string) *APIClient {
	return &APIClient{
		baseURL:   baseURL,
		authToken: authToken,
	}
}

func (c *APIClient) CallAPI(t *testing.T, endpoint string, expectedStatusCode int) interface{} {
	e := httpexpect.Default(t, c.baseURL)

	fmt.Printf("%s %s →", "GET", c.baseURL+endpoint)

	req := e.GET(endpoint).WithHeader("Authorization", "Basic "+c.authToken)

	resp := req.Expect().Status(expectedStatusCode).JSON().Raw()

	return resp
}

func (c *APIClient) ValidateResponse(t *testing.T, response interface{}, schemaFilename string) {
	schemaPath := filepath.Join("validation_schemas", schemaFilename)
	schemaContent, err := os.ReadFile(schemaPath)
	if err != nil {
		t.Fatalf("Error reading schema from file: %s", err)
	}

	schemaLoader := gojsonschema.NewStringLoader(string(schemaContent))

	var respBytes []byte
	if bytes, ok := response.([]byte); ok {
		respBytes = bytes
	} else {
		respBytes, err = json.Marshal(response)
		if err != nil {
			t.Fatalf("Error marshalling response: %s", err)
		}
	}

	documentLoader := gojsonschema.NewBytesLoader(respBytes)

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		panic(err)
	}

	if result.Valid() {
		fmt.Println(" API response is valid")
	} else {
		fmt.Println(" API response is not valid")

		var schemaMap map[string]interface{}
		err := json.Unmarshal([]byte(schemaContent), &schemaMap)
		if err != nil {
			log.Fatalf("Error parsing JSON schema: %s", err)
		}

		for _, desc := range result.Errors() {
			field := strings.TrimPrefix(desc.Field(), "data.")
			customMessage, ok := getFieldMessage(schemaMap, field)
			if ok {
				receivedValue := ""
				if match := regexp.MustCompile(`Received: (.+?)$`).FindStringSubmatch(desc.String()); len(match) > 1 {
					receivedValue = match[1]
				}
				// Replace {0} in the custom message with the received value
				customMessage = strings.Replace(customMessage, "{0}", receivedValue, 1)
				fmt.Printf("%s: %s\n", field, customMessage)
			} else {
				fmt.Printf("- %s\n", desc)
			}
		}

		var prettyJSON bytes.Buffer
		if err := json.Indent(&prettyJSON, respBytes, "", "    "); err != nil {
			t.Fatalf("failed to format JSON: %s", err)
		}

		fmt.Println("Received response:")
		fmt.Println(prettyJSON.String())
		t.Fail()
	}
}

func (c *APIClient) TestEndpoint(t *testing.T, endpoint string, expectedStatusCode int, schemaFilename string) {
	response := c.CallAPI(t, endpoint, expectedStatusCode)
	if len(schemaFilename) > 0 {
		c.ValidateResponse(t, response, schemaFilename)
	}
}

func getFieldMessage(schemaMap map[string]interface{}, field string) (string, bool) {
	properties, ok := schemaMap["properties"].(map[string]interface{})
	if !ok {
		return "", false
	}
	property, ok := properties[field].(map[string]interface{})
	if !ok {
		return "", false
	}
	message, ok := property["message"].(map[string]interface{})
	if !ok {
		return "", false
	}
	customMessage, ok := message["format"].(string)
	if !ok {
		return "", false
	}
	return customMessage, true
}
