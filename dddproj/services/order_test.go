package services

import (
	"dddproj/aggregate"

	"github.com/google/uuid"
	"testing"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Healthy boy", 60.00)
	if err != nil {
		t.Fatal(err)
	}
	peenuts, err := aggregate.NewProduct("sadg", "snacks", 2.00)
	if err != nil {
		t.Fatal(err)
	}
	wine, err := aggregate.NewProduct("wine", "wines", 2.00)
	if err != nil {
		t.Fatal(err)
	}
	return []aggregate.Product{beer, peenuts, wine}
}
func TestORder_NewOrderSeriv(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Fatal(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Fatal(err)
	}
	cust, err := aggregate.NewCustomer("jack")
	if err != nil {
		t.Fatal(err)
	}
	err = os.customers.Add(cust)
	if err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}
	err = tavern.Order(cust.GetID(), order)

	if err != nil {
		t.Fatal(err)
	}
}
