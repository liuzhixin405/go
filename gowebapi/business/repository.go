package business

import "fmt"

type (
	Employee struct {
		Id   int
		Name string
	}

	EmployeeRepository interface {
		SaveEmployee(employee Employee)
	}
	EmployeeRepositoryImpl struct {
	}
)

func (e *EmployeeRepositoryImpl) SaveEmployee(employee Employee) {
	fmt.Println("Employee saved successfully")
}
func ProvideEmployeeRepository() EmployeeRepository {
	return &EmployeeRepositoryImpl{}
}
