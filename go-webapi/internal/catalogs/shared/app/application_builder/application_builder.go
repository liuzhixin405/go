package applicationbuilder

import (
	"os"
	"path/filepath"
	"strings"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"github.com/labstack/echo/v4"
	"microproject/config"
	"microproject/internal/catalogs/shared/app/application"
	"microproject/internal/pkg/config/environment"
	"microproject/internal/pkg/constants"
)

type ApplicationBuilder struct {
	Services    *di.Builder
	Logger      *zap.SugaredLogger
	Environment environment.Environment
}

func NewApplicationBuilder(environments ...environment.Environment) *ApplicationBuilder {
	// TODO: add more services here
	env := environment.ConfigAppEnv(environments...)
	log:=createLogger()
	setConfigPath()
	builder,err:=di.NewBuilder()
	if err!=nil{
		log.Fatal(err.Error())
		return nil
	}

	return &ApplicationBuilder{
		Services:    builder,
		Logger:      log,
		Environment: env,
	}
}

func(builder *ApplicationBuilder) Build() *application.Application {
	container :=builder.Services.Build()
	echo:=container.Get("echo").(*echo.Echo)
	cfg:=container.Get("config").(*config.Config)
	return application.NewApplication(container,echo,builder.Logger,cfg)
}


func createLogger() *zap.SugaredLogger {
	logger, _:=zap.NewProduction()
	defer logger.Sync()
	log := logger.Sugar()
	return log
}

func setConfigPath(){
	wd,_:=os.Getwd()
	pn:=viper.Get(constants.PROJECT_NAME_ENV)
	if pn ==nil{
		return
	}

	for !strings.HasSuffix(wd,pn.(string)){
		wd =filepath.Dir(wd)
	}

	absCurrentDir, _ := filepath.Abs(wd)
	viper.Set(constants.AppRootPath,"config")

	configPath:=filepath.Join(absCurrentDir,"config")
	viper.Set(constants.ConfigPath,configPath)
}