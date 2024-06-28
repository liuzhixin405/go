package memory

import (
	"dddproj/aggregate"
	"dddproj/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	Customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		Customers: make(map[uuid.UUID]aggregate.Customer),
	}
}
func (mr *MemoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := mr.Customers[id]; ok {
		return customer, nil
	}
	return aggregate.Customer{}, nil
}
func (mr *MemoryRepository) Add(c aggregate.Customer) error {
	if mr.Customers == nil {
		mr.Lock()
		mr.Customers = map[uuid.UUID]aggregate.Customer{}
		mr.Unlock()
	}
	if _, ok := mr.Customers[c.GetID()]; ok {
		return fmt.Errorf("customer with id %s already exists : %w", c.GetID(), customer.ErrFailedToAddCustomer)
	}
	mr.Lock()
	mr.Customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
func (mr *MemoryRepository) Update(c aggregate.Customer) error {
	if _, ok := mr.Customers[c.GetID()]; !ok {
		return fmt.Errorf("customer with id %s not found: %w", c.GetID(), customer.ErrFailedToUpdateCustomer)
	}
	mr.Lock()
	mr.Customers[c.GetID()] = c
	mr.Unlock()
	return nil
}
