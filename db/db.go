package db

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB

func init() {
	db, err := sql.Open("postgres", "")
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)
	DB = db
	boil.SetDB(DB)
}
