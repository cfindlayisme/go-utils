package sql

import (
	"os"
	"strconv"

	"github.com/cfindlayisme/go-utils/models"
)

func GetPostgresConnectionDetails() models.PostgresConnectionDetails {
	host := os.Getenv("PG_HOST")
	portStr := os.Getenv("PG_PORT")
	username := os.Getenv("PG_USERNAME")
	password := os.Getenv("PG_PASSWORD")

	port, err := strconv.Atoi(portStr)
	if err != nil {
		port = 5432 // Default Postgres port
	}

	details := models.PostgresConnectionDetails{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
	}

	return details
}
