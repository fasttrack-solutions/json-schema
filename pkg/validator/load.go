package validator

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

//go:embed schemas/*
var schemasFileSystem embed.FS

func loadSchemas(schemasPath string, notificationTypes []string, config ValidationConfig) (map[string]gojsonschema.Schema, map[string]map[string]interface{}, error) {
	schemas := make(map[string]gojsonschema.Schema, len(notificationTypes))
	schemasRegistry := make(map[string]map[string]interface{}, len(notificationTypes))

	for _, notificationType := range notificationTypes {
		loader := gojsonschema.NewSchemaLoader()
		loader.Validate = true
		loader.Draft = gojsonschema.Draft7

		fileData, err := readSchemaFile(schemasPath, notificationType)
		if err != nil {
			return nil, nil, err
		}

		var unmarshalledSchema map[string]interface{}
		err = json.Unmarshal(fileData, &unmarshalledSchema)
		if err != nil {
			return nil, nil, fmt.Errorf("unmarshaling schema json: %v", err)
		}

		processConfig(notificationType, unmarshalledSchema, config)

		schema, err := loader.Compile(gojsonschema.NewGoLoader(unmarshalledSchema))
		if err != nil {
			return nil, nil, fmt.Errorf("compiling schema for %s: %v", notificationType, err)
		}

		schemas[notificationType] = *schema
		schemasRegistry[notificationType] = unmarshalledSchema
	}

	return schemas, schemasRegistry, nil
}

func readSchemaFile(schemasPath string, notificationType string) ([]byte, error) {
	const schemaFilePath = "%s/%s.schema.json"

	filename := fmt.Sprintf(schemaFilePath, schemasPath, notificationType)

	fileData, err := schemasFileSystem.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("reading file %s from embedded schemas file system: %v", notificationType, err)
	}

	return fileData, nil
}
