package orangesdk

import (
	"os"
)

// GetEnv reads an environment variable from the os using the key and returns
// the found value or the specified fallback value.
func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
