package env_test

import (
	"os"
	"testing"

	"github.com/cfindlayisme/go-utils/env"
)

func TestGetHttpListenPort(t *testing.T) {
	tests := []struct {
		name     string
		envPort  string
		expected int
	}{
		{"Valid PORT", "3000", 3000},
		{"Invalid PORT", "invalid", 8080},
		{"Unset PORT", "", 8080},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.envPort != "" {
				os.Setenv("PORT", test.envPort)
			} else {
				os.Unsetenv("PORT")
			}

			got := env.GetHttpListenPort()
			if got != test.expected {
				t.Errorf("GetHttpListenPort() = %d, want %d", got, test.expected)
			}
		})
	}
}
