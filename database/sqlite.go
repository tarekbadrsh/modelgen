package database

import (
	"database/sql"
)

// Sqlite driver.
type Sqlite struct {
	db *sql.DB
}

// InitDB : initialize Sqlite database object.
func (p *Sqlite) InitDB(connectionString string) error {
	db, err := InitDatabase("sqlite3", connectionString)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// DB get database object.
func (p *Sqlite) DB() *sql.DB {
	return p.db
}

// TableNamesQuery returns query to list table names in the current schema
// (not including system tables).
func (p *Sqlite) TableNamesQuery() string {
	return `SELECT name
	FROM
		sqlite_master
	WHERE
		type = 'table`
}

// ViewNamesQuery returns a list of all view names in the current schema
// (not including system views).
func (p *Sqlite) ViewNamesQuery() string {
	return `SELECT name
	FROM
		sqlite_master
	WHERE
		type = 'view`
}

// ObjectQuery returns details query table/veiw of one object.
func (p *Sqlite) ObjectQuery() string {
	return `SELECT * FROM %s LIMIT 0`
}
