package applicationbuilder

import (
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

type ApplicationBuilder struct {
	Services    *di.Builder
	Logger      *zap.SugaredLogger
	Environment environment.Environment
}

func NewApplicationBuilder(env ...environment.Environment) *ApplicationBuilder {
	// TODO: add more services here

	return &ApplicationBuilder{
		Services:    di.NewBuilder(),
		Logger:      zap.S(),
		Environment: env,
	}
}
