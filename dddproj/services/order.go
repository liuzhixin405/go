package services

import (
	"context"
	"dddproj/aggregate"
	"dddproj/domain/customer"
	"dddproj/domain/customer/memory"
	"dddproj/domain/customer/mongo"
	"dddproj/domain/product"
	prodmem "dddproj/domain/product/memory"
	"github.com/google/uuid"
	"log"
)

type OrderService struct {
	customers customer.CustomerRepository
	products  product.ProductRepository
}
type OrderConfiguration func(os *OrderService) error

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithCustomerRepository(cr customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customers = cr
		return nil
	}
}

func WithMongoCustomerRepository(ctx context.Context, connection string) OrderConfiguration {
	return func(os *OrderService) error {
		cr, err := mongo.New(ctx, connection)
		if err != nil {
			return err
		}
		os.customers = cr
		return nil
	}
}
func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		pr := prodmem.New()
		for _, p := range products {
			if err := pr.Add(p); err != nil {
				return err
			}
		}
		os.products = pr
		return nil
	}
}

func WithMemoryCustomerRepository() OrderConfiguration {
	cr := memory.New()
	return WithCustomerRepository(cr)
}

func (o *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	c, err := o.customers.Get(customerID)
	if err != nil {
		return 0, err
	}
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := o.products.GetByID(id)
		if err != nil {
			return 0, err
		}
		products = append(products, p)
		total += p.GetPrice()
	}
	log.Printf("customer: %s has ordered %d products", c.GetID(), len(products))
	return total, nil
}
