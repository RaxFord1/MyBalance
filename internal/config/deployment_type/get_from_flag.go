package deployment_type

import (
	"flag"
	"strings"
)

func GetFromFlag() DeploymentType {
	env := flag.String("env", string(Local), "environment type (e.g., local, production, etc.)")
	flag.Parse()

	if env == nil {
		return None
	}

	return DeploymentType(strings.ToLower(*env))
}
