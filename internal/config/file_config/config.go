package file_config

import (
	"MyBalance/internal/config/deployment_type"
	"MyBalance/internal/config/interface_core"
	"MyBalance/internal/config/utils"
	"MyBalance/internal/http/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	"fmt"
	"github.com/spf13/viper"
)

var _ interface_core.Cfg = (*FileConfig)(nil)

type FileConfig struct {
	DeploymentType deployment_type.DeploymentType
	Ctx            context.Context
}

func (c *FileConfig) LoadFile(ctx context.Context, envType, path string) error {

	viper.SetConfigName(envType) // Name of config file (without extension)
	viper.SetConfigType("env")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("error reading config file: %v", err.Error()))
	}

	return nil
}

func (c *FileConfig) SetDeploymentType(ctx context.Context, envType deployment_type.DeploymentType) error {

	if err := envType.Check(ctx); err != nil {
		return err
	}
	c.DeploymentType = envType

	return nil
}

func (c *FileConfig) Load(ctx context.Context, envType deployment_type.DeploymentType) error {
	c.Ctx = ctx

	if err := c.SetDeploymentType(ctx, envType); err != nil {
		return err
	}

	projectRootPath, err := FindProjectRoot(projkeys.ProjectName)
	if err != nil {
		return requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("error finding project root: %v", err.Error()))
	}
	configFolderPath := projectRootPath + "/config"

	if err = c.LoadFile(ctx, string(envType), configFolderPath); err != nil {
		return err
	}

	// Get all keys from viper
	for _, key := range viper.AllKeys() {
		newKey := utils.CheckStringAndRemovePrefix(key)
		c.Ctx.SetString(newKey, viper.GetString(key))
	}

	return nil
}

func (c *FileConfig) GetDeploymentType(ctx context.Context) deployment_type.DeploymentType {
	return c.DeploymentType
}

func NewFileConfig(ctx context.Context) *FileConfig {
	return &FileConfig{
		Ctx: ctx,
	}
}
