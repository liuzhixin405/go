package business

type (
	EmployeeService interface {
		CreateEmployee(employee Employee)
	}

	EmployeeServiceImpl struct {
		repo EmployeeRepository
	}
)

func (e *EmployeeServiceImpl) CreateEmployee(employee Employee) {
	e.repo.SaveEmployee(employee)
}

func ProvideEmployeeService(repo EmployeeRepository) EmployeeService {
	return &EmployeeServiceImpl{repo: repo}

}
