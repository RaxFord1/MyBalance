package deployment_type

import (
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"fmt"
)

type DeploymentType string

var (
	None       DeploymentType = ""
	Production DeploymentType = "production"
	Staging    DeploymentType = "staging"
	Local      DeploymentType = "local"
)

func (d DeploymentType) Check(ctx context.Context) error {
	switch d {
	case Production, Staging, Local:
		return nil
	default:
		return requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("unknown deployment type: %s", d))
	}
}
