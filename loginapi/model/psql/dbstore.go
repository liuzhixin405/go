package psql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	db, err := sql.Open("mysql", "root:1230@tcp(192.168.237.240:3306)/test")
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}
	return store, nil
}

func (s *Store) Close() {
	s.db.Close()
}
