package services

import (
	"dddproj/aggregate"
	"github.com/google/uuid"
	"testing"
)

func Test_Tavern(t *testing.T) {
	products := init_products(t)
	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
		//WithMongoCustomerRepository(context.Background(),"mongo connection string")
	)
	if err != nil {
		t.Error(err)
	}
	tavern, err := NewTavern(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}
	cust, err := aggregate.NewCustomer("jack")
	if err != nil {
		t.Error(err)
	}
	if err = os.customers.Add(cust); err != nil {
		t.Fatal(err)
	}
	order := []uuid.UUID{
		products[0].GetID(),
	}

	err = tavern.Order(cust.GetID(), order)
	if err != nil {
		t.Error(err)
	}

}
