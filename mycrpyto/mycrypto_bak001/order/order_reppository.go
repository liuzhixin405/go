package order

import "gorm.io/gorm"

type OrderRepository interface {
	Create(order *Order) (int, error)
	UpdateStatus(orderID int, status string) error
}

type MySqlUserRepository struct {
	db *gorm.DB
}

func NewMySqlUserRepository(db *gorm.DB) *MySqlUserRepository {
	return &MySqlUserRepository{db: db}
}

func (repo *MySqlUserRepository) Create(order *Order) (int, error) {
	repo.db.Create(order)
	// insert into orders (order_id, status) values (?,?)
	// returning id
	return 1, nil
}

func (repo *MySqlUserRepository) UpdateStatus(orderID int, status string) error {
	// update orders set status =? where order_id =?
	var order Order
	repo.db.Model(&order).Where("order_id = ?", orderID).Update("status", status)
	return nil
}
