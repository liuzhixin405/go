package aggregate

import (
	"dddproj/entity"
	"dddproj/valueobject"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("customer has invalid name ")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		Name: name,
		ID:   uuid.New(),
	}
	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueobject.Transaction, 0),
	}, nil
}
func (c *Customer) GetID() uuid.UUID {
	return c.person.ID
}
func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.ID = id
}
func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	c.person.Name = name
}
func (c *Customer) GetName() string {
	return c.person.Name
}
