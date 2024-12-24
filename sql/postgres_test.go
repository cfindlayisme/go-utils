package sql_test

import (
	"os"
	"testing"

	"github.com/cfindlayisme/go-utils/models"
	"github.com/cfindlayisme/go-utils/sql"
)

func TestGetPostgresConnectionDetails(t *testing.T) {
	tests := []struct {
		envHost     string
		envPort     string
		envUsername string
		envPassword string
		expected    models.PostgresConnectionDetails
	}{
		{"localhost", "5432", "user", "password", models.PostgresConnectionDetails{"localhost", 5432, "user", "password"}},
		{"localhost", "invalid", "user", "password", models.PostgresConnectionDetails{"localhost", 5432, "user", "password"}},
		{"localhost", "", "user", "password", models.PostgresConnectionDetails{"localhost", 5432, "user", "password"}},
	}

	for _, test := range tests {
		os.Setenv("PG_HOST", test.envHost)
		os.Setenv("PG_PORT", test.envPort)
		os.Setenv("PG_USERNAME", test.envUsername)
		os.Setenv("PG_PASSWORD", test.envPassword)

		details := sql.GetPostgresConnectionDetails()

		if details.Host != test.expected.Host {
			t.Errorf("expected Host to be '%s', got '%s'", test.expected.Host, details.Host)
		}
		if details.Port != test.expected.Port {
			t.Errorf("expected Port to be %d, got %d", test.expected.Port, details.Port)
		}
		if details.Username != test.expected.Username {
			t.Errorf("expected Username to be '%s', got '%s'", test.expected.Username, details.Username)
		}
		if details.Password != test.expected.Password {
			t.Errorf("expected Password to be '%s', got '%s'", test.expected.Password, details.Password)
		}
	}
}

func TestConstructPostgresConnectionString(t *testing.T) {
	tests := []struct {
		details  models.PostgresConnectionDetails
		database string
		expected string
	}{
		{models.PostgresConnectionDetails{"localhost", 5432, "user", "password"}, "testdb", "postgres://user:password@localhost:5432/testdb"},
		{models.PostgresConnectionDetails{"localhost", 5432, "user", "password"}, "", "postgres://user:password@localhost:5432/"},
	}

	for _, test := range tests {
		connectionString := sql.ConstructPostgresConnectionString(test.details, test.database)
		if connectionString != test.expected {
			t.Errorf("expected connection string to be '%s', got '%s'", test.expected, connectionString)
		}
	}
}
