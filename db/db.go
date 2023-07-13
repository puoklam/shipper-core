package db

import (
	"database/sql"
	"os"
	"sync"
	"time"

	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

var DB *sql.DB
var once sync.Once

func Init() (err error) {
	once.Do(func() {
		DB, err = sql.Open("postgres", os.Getenv("DB_URL"))
		if err != nil {
			return
		}
		DB.SetMaxOpenConns(25)
		DB.SetMaxIdleConns(25)
		DB.SetConnMaxLifetime(5 * time.Minute)
		boil.SetDB(DB)
	})
	return
}
