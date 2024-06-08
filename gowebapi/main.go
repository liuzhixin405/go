package main

import (
	applicationbuilder "gowebapi/application_builder"
	"gowebapi/routes"
)

func main() {
	builder := applicationbuilder.NewApplicationBuilder()
	builder.AddRoutes(routes.TemperatureRoutes{})
	app := builder.Build()
	app.Run("localhost:8088")
}
