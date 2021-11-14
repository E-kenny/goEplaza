package database

import (
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Connection() (*sqlx.DB, error) {
	//capture connection properties
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   os.Getenv("HOST"),
		DBName: os.Getenv("DBNAME"),
	}

	//Get database handle

	db, err := sqlx.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	return db, nil
}
