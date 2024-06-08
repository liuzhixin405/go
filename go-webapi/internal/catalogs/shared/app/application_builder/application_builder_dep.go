package applicationbuilder

import (
	"microproject/config"
	config2 "microproject/internal/pkg/config"
	"microproject/internal/pkg/config/environment"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
)

func (b *ApplicationBuilder) AddCore() {
	legDep := di.Def{
		Name:  "zap",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return b.Logger, nil
		},
	}

	configDep := di.Def{
		Name:  "config",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			env := ctn.Get("env").(environment.Environment)
			return config.NewConfig(env)
		},
	}

	err := config2.AddEnv(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(legDep)
	if err != nil {
		b.Logger.Fatal(err)
	}

	err = b.Services.Add(configDep)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddInfrastructure() {
	err := addEcho(b.Services)
	if err != nil {
		b.Logger.Fatal(err)
	}
}

func (b *ApplicationBuilder) AddRoutes() {

	routesDep := di.Def{
		Name:  "routes",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {

			return nil, nil
		},
	}
	err := b.Services.Add(routesDep)
	if err != nil {
		b.Logger.Fatal(err)
	}

}
func (b *ApplicationBuilder) AddRepositories() {

}

func addEcho(container *di.Builder) error {
	echoDep := di.Def{
		Name:  "echo",
		Scope: di.App,
		Build: func(ctn di.Container) (interface{}, error) {
			return echo.New(), nil
		},
	}

	err := container.Add(echoDep)
	if err != nil {
		return err
	}
	return nil
}
