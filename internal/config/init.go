package config

import (
	"MyBalance/internal/config/deployment_type"
	"MyBalance/internal/config/file_config"
	"MyBalance/internal/config/interface_core"
	"MyBalance/internal/config/memory_config"
	"MyBalance/internal/context"
)

var config interface_core.Cfg

func Init(ctx context.Context) error {
	switch configSourceType {
	case TypeOfConfigNotDefined:
		config = file_config.NewFileConfig(ctx)
	case TypeOfConfigFromFile:
		config = file_config.NewFileConfig(ctx)
	case TypeOfConfigFromMemory:
		config = memory_config.NewMemoryConfig(ctx)
	}

	var deploymentType deployment_type.DeploymentType
	switch deploymentInfoSourceType {
	case DeploymentInfoNotDefined:
		deploymentType = deployment_type.GetFromFlag()
	case DeploymentInfoFromFlag:
		deploymentType = deployment_type.GetFromFlag()
	case DeploymentInfoFromEnv:
		deploymentType = deployment_type.GetFromEnv()
	}

	if deploymentType == deployment_type.None {
		deploymentType = deployment_type.Local
	}

	if err := deploymentType.Check(ctx); err != nil {
		return err
	}

	err := config.Load(ctx, deploymentType)
	if err != nil {
		return err
	}

	return nil
}
