package user

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *User) (int, error)
	FindByID(id int) (User, error)
	UpdateBalance(userID int, amount float64) error
	Login(email string, password string) (bool, error)
}

type MySqlUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *MySqlUserRepository {
	return &MySqlUserRepository{db: db}
}

func (repo *MySqlUserRepository) Login(email string, password string) (bool, error) {
	var user User
	err := repo.db.Where("email = ? AND password = ?", email, password).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (repo *MySqlUserRepository) Create(user *User) (int, error) {
	err := repo.db.Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

func (repo *MySqlUserRepository) FindByID(id int) (User, error) {

	var user User
	err := repo.db.First(&user, id).Error
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (repo *MySqlUserRepository) UpdateBalance(userID int, amount float64) error {
	err := repo.db.Model(&User{}).Where("id = ?", userID).Update("balance", gorm.Expr("balance + ?", amount)).Error
	if err != nil {
		return err
	}
	return nil
}
