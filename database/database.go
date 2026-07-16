package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

var db *sqlx.DB

const SCHEMA = `
PRAGMA journal_mode = WAL;
PRAGMA busy_timeout = 5000;

CREATE TABLE IF NOT EXISTS archers (
	archer_id INTEGER PRIMARY KEY,
	first_name TEXT NOT NULL COLLATE NOCASE,
	last_name TEXT NOT NULL COLLATE NOCASE,
	school TEXT,
	bowtype TEXT NOT NULL,
	age INTEGER NOT NULL
)
	
CREATE TABLE IF NOT EXISTS events (
	event_id INTEGER PRIMARY KEY,
	event_name TEXT NOT NULL COLLATE NOCASE,
	date DATETIME NOT NULL,
	round_num INTEGER NOT NULL,
	session INTEGER CHECK (round_num == 1 OR round_num == 2)
);
`

func OpenDatabase(dbPath string) error {
	var err error
	db, err = sqlx.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	db.MustExec(SCHEMA)
	return nil
}

func CloseDatabase() {
	if db != nil {
		db.Close()
	}
}
