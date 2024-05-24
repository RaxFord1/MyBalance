package memory_config

import (
	"MyBalance/internal/config/deployment_type"
	"MyBalance/internal/config/interface_core"
	"MyBalance/internal/context"
)

var _ interface_core.Cfg = (*MemoryConfig)(nil)

type MemoryConfig struct {
	DeploymentType deployment_type.DeploymentType
	Ctx            context.Context
}

func (m *MemoryConfig) Load(ctx context.Context, envType deployment_type.DeploymentType) error {
	envVars := loadEnvVars()

	for key, value := range envVars {
		ctx.SetString(key, value)
	}

	return nil
}

func (m *MemoryConfig) SetDeploymentType(ctx context.Context, envType deployment_type.DeploymentType) error {
	m.DeploymentType = envType
	return nil
}

func (m *MemoryConfig) GetDeploymentType(ctx context.Context) deployment_type.DeploymentType {
	return m.DeploymentType
}

func NewMemoryConfig(ctx context.Context) *MemoryConfig {
	return &MemoryConfig{
		Ctx: ctx,
	}
}
