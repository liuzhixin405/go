package config

import (
	"microproject/internal/pkg/config/environment"

	"github.com/sarulabs/di"
)

func AddEnv(container *di.Builder) error {
	envDep := di.Def{
		Name:  "env",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return environment.ConfigAppEnv(), nil
		},
	}
	return container.Add(envDep)
}
