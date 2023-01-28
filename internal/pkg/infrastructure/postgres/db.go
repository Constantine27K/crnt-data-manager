package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Options struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func NewPostgres(opts Options) (*sql.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		opts.Host, opts.Port, opts.User, opts.Password, opts.DBName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
