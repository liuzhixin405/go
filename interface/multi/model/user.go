package model

import (
	"context"
	"errors"
	"fmt"
)

type Db struct {
}

func New() *Db {
	return &Db{}
}

type User struct {
	ID    string
	Score int
}

func (d *Db) GetUser(ctx context.Context, id string) (*User, error) {
	return &User{
		ID:    id,
		Score: 7,
	}, nil
}

func (d *Db) CreateUser(ctx context.Context, id string) error {
	return nil
}
func (d *Db) GetUserScore(ctx context.Context, id string) (int, error) {
	return 7, nil
}
func (d *Db) DeleteUser(ctx context.Context, id string) error {
	fmt.Println("delete Successful")
	return nil
}

func (d *Db) GetTopUsers(ctx context.Context, count int) ([]*User, error) {
	return nil, errors.New("Notimplemented,Evertras is lazy")
}

func (d *Db) AwardPoints(ctx context.Context, ids []string, score int) error {
	return nil
}
