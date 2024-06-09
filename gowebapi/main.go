package main

import (
	applicationbuilder "gowebapi/application_builder"
	"gowebapi/business"
	"gowebapi/routes"
)

func main() {

	employeeService := InitializeService(InitializeRepository())
	employeeService.CreateEmployee(business.Employee{Name: "John", Id: 30})

	builder := applicationbuilder.NewApplicationBuilder()
	builder.AddRoutes(routes.TemperatureRoutes{})
	app := builder.Build()
	app.Run("localhost:8088")
}
