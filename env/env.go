package env

import (
	"os"
	"strconv"
)

func GetHttpListenPort() int {
	if os.Getenv("PORT") != "" {
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			return 8080
		}
		return port
	}
	return 8080
}
