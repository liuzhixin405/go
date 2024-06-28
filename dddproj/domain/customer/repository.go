package customer

import (
	"dddproj/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound       = errors.New("customer not found int the repository")
	ErrFailedToAddCustomer    = errors.New("failed to add customer int the repository")
	ErrFailedToUpdateCustomer = errors.New("failed to update customer int the repository")
	ErrFailedToDeleteCustomer = errors.New("failed to delete customer int the repository")
)

type CustomerRepository interface {
	Get(uuid uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
	//Delete(uuid.UUID) error
}
