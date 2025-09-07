package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db   *sql.DB
	once sync.Once
)

const dsn =  "root:Ygwatchguard@#123@tcp(localhost:3306)/tasknest_db?parseTime=true"


func GetDbInstance() (*sql.DB, error) {
	var err error

	once.Do(func() {

		db, err = sql.Open("mysql", dsn)
		if err != nil {
			log.Fatalf("Error opening DB: %v", err)
		}

		// Connection Pool Settings
		db.SetConnMaxIdleTime(25 * time.Second)
		db.SetConnMaxLifetime(time.Hour)
		db.SetMaxIdleConns(50)
		db.SetMaxOpenConns(100)

		// Test connection
		err = db.Ping()
		if err != nil {
			log.Fatalf("Error connng to DB: %v", err)
		}

		fmt.Println("Successfully connected to MySQL Database!")
	})

	return db, err
}
