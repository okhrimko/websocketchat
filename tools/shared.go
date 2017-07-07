package shared

import (
	"os"
)

func Getenv(key string) string {
	port := os.Getenv(key)
	if len(port) == 0 {
		return "8080"
	}
	return port
}
