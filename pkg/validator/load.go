package validator

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/*
var schemasFileSystem embed.FS

const (
	schemasPathRealTimeEvents = "schemas/real-time-events"
	schemasPathOperatorAPI    = "schemas/operator-api/get"
)

func loadSchemas(schemasPath string, notificationTypes []string) (map[string]gojsonschema.JSONLoader, map[string]map[string]interface{}, error) {
	eventSchemas := make(map[string]gojsonschema.JSONLoader, len(notificationTypes))
	publicEventSchemas := make(map[string]map[string]interface{}, len(notificationTypes))

	for _, notificationType := range notificationTypes {
		filename := fmt.Sprintf("%s/%s.schema.json", schemasPath, notificationType)

		fileData, err := schemasFileSystem.ReadFile(filename)
		if err != nil {
			return nil, nil, fmt.Errorf("reading file %s from embedded schemas file system: %v", notificationType, err)
		}

		eventSchemas[notificationType] = gojsonschema.NewBytesLoader(fileData)

		var rawSchema map[string]interface{}
		err = json.Unmarshal(fileData, &rawSchema)
		if err != nil {
			return nil, nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		publicEventSchemas[notificationType] = rawSchema
	}

	return eventSchemas, publicEventSchemas, nil
}
