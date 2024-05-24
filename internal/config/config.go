package config

import (
	"MyBalance/internal/context"
	"MyBalance/internal/http/requesto"
	"MyBalance/internal/projkeys"
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Cfg interface {
	Load(ctx context.Context, envType DeploymentType) error
	SetDeploymentType(ctx context.Context, envType DeploymentType) error
	GetDeploymentType(ctx context.Context) DeploymentType
}

type Config struct {
	DeploymentType DeploymentType
	Ctx            context.Context
}

func (c *Config) LoadFile(ctx context.Context, envType, path string) error {

	viper.SetConfigName(envType) // Name of config file (without extension)
	viper.SetConfigType("env")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return requesto.InternalError.NewWithMsg(ctx, fmt.Sprintf("error reading config file: %v", err.Error()))
	}

	return nil
}

func (c *Config) SetDeploymentType(ctx context.Context, envType DeploymentType) error {

	if err := envType.Check(ctx); err != nil {
		return err
	}
	c.DeploymentType = envType

	return nil
}

func checkStringAndRemovePrefix(key string) string {
	if strings.HasPrefix(key, projkeys.ProjectName+".") {
		// remove prefix
		if key2 := strings.Replace(key, projkeys.ProjectName+".", "", 1); key2 != "" {
			return key2
		} else {
			return key
		}
	} else if strings.HasPrefix(key, strings.ToLower(projkeys.ProjectName)+".") {
		// remove prefix
		if key2 := strings.Replace(key, strings.ToLower(projkeys.ProjectName)+".", "", 1); key2 != "" {
			return key2
		} else {
			return key
		}
	}

	return key
}

func (c *Config) Load(ctx context.Context, envType DeploymentType) error {
	c.Ctx = ctx

	if err := config.SetDeploymentType(ctx, envType); err != nil {
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
		newKey := checkStringAndRemovePrefix(key)
		c.Ctx.SetString(newKey, viper.GetString(key))
	}

	return nil
}

func (c *Config) GetDeploymentType(ctx context.Context) DeploymentType {
	return c.DeploymentType
}

func GetDeploymentType(ctx context.Context) DeploymentType {
	return config.GetDeploymentType(ctx)
}
