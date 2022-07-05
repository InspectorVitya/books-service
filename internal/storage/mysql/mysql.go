package mysql

import (
	"fmt"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

type StorageMySQL struct {
	db *sqlx.DB
}

func New(host, nameDB, user, psw, port string) (*StorageMySQL, error) {
	fmt.Println(getDBURL(host, nameDB, user, psw, port))
	db, err := sqlx.Connect("mysql", getDBURL(host, nameDB, user, psw, port))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &StorageMySQL{db: db}, err
}

func (s *StorageMySQL) Stop() error {
	return s.db.Close()
}

func getDBURL(host, nameDB, user, psw, port string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, psw, host, port, nameDB)
}
