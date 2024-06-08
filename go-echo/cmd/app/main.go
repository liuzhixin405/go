package main

import applicationbuilder "microproject/internal/catalogs/shared/app/application_builder"

func main() {
	builder := applicationbuilder.NewApplicationBuilder()
	builder.AddCore()
	builder.AddInfrastructure()
	app := builder.Build()
	app.Run()
}
