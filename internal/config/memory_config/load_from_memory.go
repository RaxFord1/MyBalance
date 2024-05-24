package memory_config

import (
	"os"
	"strings"
)

// loadEnvVars loads all environment variables into a map
func loadEnvVars() map[string]string {
	envVars := make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.SplitN(env, "=", 2)
		if len(pair) == 2 {
			envVars[pair[0]] = pair[1]
		}
	}

	return envVars
}
