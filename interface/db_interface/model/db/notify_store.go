package db

import "fmt"

func (s *Store) Notify(message string) {
	fmt.Println(message)
}
