package test

import (
	"testing"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestVerifyPathExists(t *testing.T) {
	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/sys/bus/w1/devices/device1", []byte{}, 0644)

	tests := []struct {
		name        string
		path        string
		setupFs     func()
		expected    string
		expectError bool
	}{
		{
			name: "Path exists",
			path: "/sys/bus/w1/devices/",
			setupFs: func() {
				fs.MkdirAll("/sys/bus/w1/devices", 0755)
				afero.WriteFile(fs, "/sys/bus/w1/devices/device1", []byte{}, 0644)
			},
			expected:    "device1",
			expectError: false,
		},
		{
			name: "Path does not exist",
			path: "/sys/bus/w1/devices_nonexistent/",
			setupFs: func() {},
			expected:    "error",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setupFs()
			name, err := VerifyPathExists(fs, tt.path)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.expected, name)
		})
	}
}

// Updated function to use afero.Fs for filesystem operations
func VerifyPathExists(fs afero.Fs, path string) (string, error) {
	entries, err := afero.ReadDir(fs, path)
	if err != nil {
		return "error", err
	}

	for _, e := range entries {
		return e.Name(), nil
	}
	return "error", nil
}