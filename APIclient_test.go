package json_schema

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReadSchemaFile(t *testing.T) {
	content, err := readSchemaFile("operator-api/get/user_details.schema.json")
	require.NoError(t, err)
	require.NotEmpty(t, content)
}

func TestReadSchemaFileRejectsPathsOutsideSchemasFolder(t *testing.T) {
	for _, name := range []string{
		"../../../go.mod",
		"operator-api/../../../../go.mod",
		"/etc/passwd",
	} {
		_, err := readSchemaFile(name)
		require.Error(t, err, name)
	}
}
