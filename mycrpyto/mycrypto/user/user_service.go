package user

import (
	"errors"
	"log"
)

type UserService interface {
	Register(username, email, password string) (int, error)

	Login(username, password string) (bool, error)

	Deposit(userID int, amount float64) error
	//GetBlance(userID int) (float64, error)
	GetUserInfo(userID int) (User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Login(username, password string) (bool, error) {
	isLogin, err := s.repo.Login(username, password)
	if err != nil {
		return false, err
	}
	return isLogin, nil
}

func (s *userService) GetUserInfo(userID int) (User, error) {
	user, err := s.repo.FindByID(userID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

//	func (s *userService) GetBlance(userID int) (float64, error) {
//		balance, err := s.repo.GetBalance(userID)
//		if err != nil {
//			return 0, err
//		}
//		return balance, nil
//	}
func (s *userService) Register(username, email, password string) (int, error) {
	if username == "" || email == "" || password == "" {
		return 0, errors.New("username, email and password are required")
	}
	user := &User{
		Username: username,
		Email:    email,
		Password: password,
		Balance:  1000,
	}
	id, err := s.repo.Create(user)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (s *userService) Deposit(userID int, amount float64) error {
	err := s.repo.UpdateBalance(userID, amount)
	if err != nil {
		log.Printf("Failed to deposit amount: %v", err)
		return err
	}
	return err
}
