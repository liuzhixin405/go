package main

import applicationbuilder "microproject/internal/catalogs/shared/app/application_builder"
func main() {
	builder := applicationbuilder.NewApplicationBuilder()
	app:= builder.Build()
	app.Run()
}
