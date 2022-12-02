package crud

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Welcome1@tcp(localhost:3306)/hrd")

	if err != nil {
		return nil, err
	}

	return db, nil
}
