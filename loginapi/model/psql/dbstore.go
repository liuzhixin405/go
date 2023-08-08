package psql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type Store struct {
	db *sql.DB
}

func NewStore() (*Store, error) {
	db, err := sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/mydb")
	if err != nil {
		return nil, err
	}

	store := &Store{db: db}
	return store, nil
}

func (s *Store) Close() {
	s.db.Close()
}
