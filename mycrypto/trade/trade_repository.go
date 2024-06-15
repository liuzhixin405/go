package trade

import "gorm.io/gorm"

type TradeRepository interface {
	SaveTrade(trade *Trade) (int, error)
}

type MySqlUserRepository struct {
	db *gorm.DB
}

func NewMySqlUserRepository(db *gorm.DB) *MySqlUserRepository {
	return &MySqlUserRepository{db: db}
}

func (repo *MySqlUserRepository) SaveTrade(trade *Trade) (int, error) {
	// TODO: implement saving trade to database
	repo.db.Create(trade)
	return 0, nil
}
