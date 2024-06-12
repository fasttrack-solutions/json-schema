package validator

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/*
var schemasFileSystem embed.FS

func loadSchemas(schemasPath string, notificationTypes []string) (map[string]gojsonschema.JSONLoader, map[string]map[string]interface{}, error) {
	schemas := make(map[string]gojsonschema.JSONLoader, len(notificationTypes))
	schemasRegistry := make(map[string]map[string]interface{}, len(notificationTypes))

	for _, notificationType := range notificationTypes {
		fileData, err := readSchemaFile(schemasPath, notificationType)
		if err != nil {
			return nil, nil, err
		}

		schemas[notificationType] = gojsonschema.NewBytesLoader(fileData)

		var unmarshalledSchema map[string]interface{}
		err = json.Unmarshal(fileData, &unmarshalledSchema)
		if err != nil {
			return nil, nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		schemasRegistry[notificationType] = unmarshalledSchema
	}

	return schemas, schemasRegistry, nil
}

func readSchemaFile(schemasPath string, notificationType string) ([]byte, error) {
	filename := fmt.Sprintf("%s/%s.schema.json", schemasPath, notificationType)

	fileData, err := schemasFileSystem.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file %s from embedded schemas file system: %v", notificationType, err)
	}

	return fileData, nil
}
