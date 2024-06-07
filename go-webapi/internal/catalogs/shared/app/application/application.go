package application

import (
	"context"
	"errors"
	"log"
	"microproject/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/sarulabs/di"
	"go.uber.org/zap"
)

type Application struct {
	Contrainer di.Container
	Echo       *echo.Echo
	Logger     *zap.SugaredLogger
	Cfg        *config.Config
}

func NewApplication(container di.Container, echo *echo.Echo, logger *zap.SugaredLogger, cfg *config.Config) *Application {
	return &Application{
		Contrainer: container,
		Echo:       echo,
		Logger:     logger,
		Cfg:        cfg,
	}
}
func (app *Application) Run() {
	defaultDuration := time.Second * 20
	startCtx, cancel := context.WithTimeout(context.Background(), defaultDuration)
	defer cancel()
	app.Start(startCtx)
	<-app.Wait()
	stopCtx, cancel := context.WithTimeout(context.Background(), defaultDuration)
	defer cancel()
	app.Stop(stopCtx)

}

func (app *Application) Start(ctx context.Context) {
	app.Logger.Info("Starting application")
	echoStartHook(ctx, app)
}

func (app *Application) Stop(ctx context.Context) {
	app.Logger.Info("Stopping application")
	echoStartHook(ctx, app)
	app.Contrainer.Delete()
}

func (app *Application) Wait() <-chan os.Signal {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	return sigChan
}

func echoStartHook(startCtx context.Context, app *Application) {
	if err := app.Echo.Start(app.Cfg.EchohttpOptions.Port); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start echo server: %v", err)
	}
	log.Println("Stopped serving new connections")
}
