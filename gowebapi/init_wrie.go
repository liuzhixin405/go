//go:build wireinject
// +build wireinject

package main

import (
	"gowebapi/business"

	"github.com/google/wire"
)

func InitializeRepository() business.EmployeeRepository {
	wire.Build(business.ProvideEmployeeRepository)
	return nil
}

func InitializeService(repo business.EmployeeRepository) business.EmployeeService {
	wire.Build(business.ProvideEmployeeService)
	return nil
}
