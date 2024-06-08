package applicationbuilder

import (
	"gowebapi/application"

	"github.com/gin-gonic/gin"
)

type ApplicationBuilder struct {
	Router *gin.Engine
}
type RouteConfigurator interface {
	ConfigureRoutes(router *gin.Engine)
}

func NewApplicationBuilder() *ApplicationBuilder {
	router := gin.Default()
	return &ApplicationBuilder{Router: router}
}

func (builder *ApplicationBuilder) AddRoutes(configurators ...RouteConfigurator) {
	for _, configurator := range configurators {
		configurator.ConfigureRoutes(builder.Router)
	}
	builder.Router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})

	})
}

func (builder *ApplicationBuilder) Build() *application.Application {
	return &application.Application{
		Router: builder.Router,
	}
}
