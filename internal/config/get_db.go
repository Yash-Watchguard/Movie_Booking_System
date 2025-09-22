package config

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func GetDbInstance() (*sql.DB, error) {
	var err error

	once.Do(func() {
		
		db, err = sql.Open("sqlite3", "./bookingsystem.db")
		if err != nil {
			log.Fatalf("Error opening DB: %v", err)
		}

		db.SetMaxOpenConns(1) 
		db.SetMaxIdleConns(1)

		// Test connection
		err = db.Ping()
		if err != nil {
			log.Fatalf("Error connecting to DB: %v", err)
		}

		fmt.Println("Successfully connected to SQLite Database!")

		
		_, err = db.Exec("PRAGMA foreign_keys = ON;")
		if err != nil {
			log.Fatalf("Error enabling foreign keys: %v", err)
		}

		
		err = initSchema(db)
		if err != nil {
			log.Fatalf("Error creating schema: %v", err)
		}
	})

	return db, err
}

func initSchema(db *sql.DB) error {
	schema := `
CREATE TABLE IF NOT EXISTS users (
user_id TEXT PRIMARY KEY,
name TEXT NOT NULL,
email TEXT NOT NULL UNIQUE,
phone_number TEXT NOT NULL UNIQUE,
password TEXT NOT NULL,
role TEXT NOT NULL CHECK(role IN ('ADMIN', 'CUSTOMER'))
);
CREATE TABLE IF NOT EXISTS movies (
    movie_id TEXT PRIMARY KEY,
    movie_name TEXT NOT NULL,
    movie_type TEXT NOT NULL,
    duration INTEGER NOT NULL
    );
 
CREATE TABLE IF NOT EXISTS shows (
show_id TEXT PRIMARY KEY,
movie_id TEXT NOT NULL,
start_time DATETIME NOT NULL,
end_time DATETIME NOT NULL,
total_seats INTEGER NOT NULL,
  available_seats INTEGER NOT NULL,
  FOREIGN KEY(movie_id) REFERENCES movies(movie_id) ON DELETE CASCADE
);
 
CREATE TABLE IF NOT EXISTS tickets (
ticket_id TEXT PRIMARY KEY,
show_id TEXT NOT NULL,
user_id TEXT NOT NULL,
booking_time DATETIME NOT NULL,
FOREIGN KEY(show_id) REFERENCES shows(show_id) ON DELETE CASCADE,
FOREIGN KEY(user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
`

	_, err := db.Exec(schema)
	return err
}
