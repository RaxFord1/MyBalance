package deployment_type

import (
	"MyBalance/internal/projkeys"
	"os"
	"strings"
)

func GetFromEnv() DeploymentType {
	env := os.Getenv(projkeys.DeploymentType)

	if env == "" {
		return None
	}

	return DeploymentType(strings.ToLower(env))
}
