package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"os"
)

func Connect(dbName, login, password, url, port string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", login, password, url, port, dbName))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	migrations, err := os.ReadFile("./migrations/0000001_create_all.up.sql")
	if err != nil {
		return nil, err
	}

	fmt.Println(string(migrations))
	if _, err := db.Exec(string(migrations)); err != nil {
		return nil, err
	}

	return db, nil
}
