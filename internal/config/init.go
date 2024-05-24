package config

import (
	"MyBalance/internal/context"
	"flag"
)

var config Cfg

var (
	_ Cfg = (*Config)(nil)
)

func Init(ctx context.Context) error {
	config = &Config{}

	env := flag.String("env", string(Local), "environment type (e.g., local, production, etc.)")
	flag.Parse()

	if env == nil {
		local := string(Local)
		env = &local
	}

	deploymentType := DeploymentType(*env)

	if err := deploymentType.Check(ctx); err != nil {
		return err
	}

	err := config.Load(ctx, deploymentType)
	if err != nil {
		return err
	}

	return nil
}
