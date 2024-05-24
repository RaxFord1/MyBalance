package config

import (
	"MyBalance/internal/config/deployment_type"
	"MyBalance/internal/context"
	"MyBalance/internal/http/requesto"
	"fmt"
)

// T config type - where we get info about our env (from file, from memory, etc.)
type T int

const (
	TypeOfConfigNotDefined T = iota
	TypeOfConfigFromMemory
	TypeOfConfigFromFile
)

func (r T) Check(ctx context.Context) error {
	switch r {
	case TypeOfConfigNotDefined, TypeOfConfigFromMemory, TypeOfConfigFromFile:
		return nil
	default:
		return requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("invalid config type"))
	}
}

// di deployment info type - where we get info about what are we (local/production/etc.)
type di int

const (
	DeploymentInfoNotDefined di = iota
	DeploymentInfoFromFlag
	DeploymentInfoFromEnv
)

var configSourceType T
var deploymentInfoSourceType di

func SetConfigType(ctx context.Context, envType T) {
	configSourceType = envType
}

func SetDeploymentInfoSource(ctx context.Context, envType di) {
	deploymentInfoSourceType = envType
}

func GetDeploymentType(ctx context.Context) deployment_type.DeploymentType {
	return config.GetDeploymentType(ctx)
}
