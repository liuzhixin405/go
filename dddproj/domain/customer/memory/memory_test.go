package memory_test

import (
	"dddproj/aggregate"
	"dddproj/domain/customer"
	"dddproj/domain/customer/memory"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestMemoryRepository_GetCustomer(t *testing.T) {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	cust, err := aggregate.NewCustomer("jack")
	if err != nil {
		t.Fatal(err)
	}
	id := cust.GetID()
	repo := memory.MemoryRepository{
		Customers: map[uuid.UUID]aggregate.Customer{
			id: cust,
		},
	}
	testCases := []testCase{
		{
			name:        "no customer by id",
			id:          uuid.MustParse("xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"),
			expectedErr: customer.ErrCustomerNotFound,
		}, {
			name:        "customer by id",
			id:          id,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedErr) {
				t.Errorf("excepted: %v, got: %v", tc.expectedErr, err)
			}
		})
	}
}
