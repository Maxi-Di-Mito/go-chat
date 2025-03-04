package main

import (
	"database/sql"
	"log"
)

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./app.db")
	if err != nil {
		log.Fatal(err)
	}

	stmt := `
  CREATE TABLE IF NOT EXISTS prices (
  id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
  price INTEGER,
  name TEXT
  );`

	_, err = DB.Exec(stmt)
	if err != nil {
		log.Fatalf("Error creating table %q: %s\n", err, stmt)
	}

}

func GetPrices() {
	stmt := ` SELECT * FROM prices`

	rows, err := DB.Exec(stmt)
	if err != nil {
		log.Fatalf("Error reading table %q: %s", err, stmt)
	}
}
