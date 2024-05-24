package interface_core

import (
	"MyBalance/internal/config/deployment_type"
	"MyBalance/internal/context"
)

type Cfg interface {
	Load(ctx context.Context, envType deployment_type.DeploymentType) error
	SetDeploymentType(ctx context.Context, envType deployment_type.DeploymentType) error
	GetDeploymentType(ctx context.Context) deployment_type.DeploymentType
}
