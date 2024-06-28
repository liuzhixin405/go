package product

import (
	"dddproj/aggregate"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound = errors.New("product not found")
	ErrProductAdd      = errors.New("product add error")
	ErrProductUpdate   = errors.New("product update error")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
