package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_loadSchemas(t *testing.T) {
	tests := []struct {
		name              string
		schemasPath       string
		notificationTypes []string
		expectError       bool
	}{
		{
			name:              "Valid schemas",
			schemasPath:       pathSchemasRealTime,
			notificationTypes: []string{"login_v2", "login_v2"},
			expectError:       false,
		},
		{
			name:              "Empty schemas",
			schemasPath:       pathSchemasRealTime,
			notificationTypes: []string{},
			expectError:       false,
		},
		{
			name:              "Invalid schemasPath",
			schemasPath:       "/invalid/path",
			notificationTypes: []string{"login_v2", "login_v2"},
			expectError:       true,
		},
		{
			name:              "Invalid notificationTypes",
			schemasPath:       "/path/to/schemas",
			notificationTypes: []string{"type1", "type2", "type3"},
			expectError:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualSchemas, actualRegistry, err := loadSchemas(tt.schemasPath, tt.notificationTypes, ValidationConfig{})

			if tt.expectError {
				require.Error(t, err)
				require.Nil(t, actualSchemas)
				require.Nil(t, actualRegistry)
				return
			}

			require.NoError(t, err)
			require.NotNil(t, actualSchemas)
			require.NotNil(t, actualRegistry)
		})
	}
}
