package database

import (
	"database/sql"
)

// Mysql driver.
type Mysql struct {
	db *sql.DB
}

// InitDB : initialize Mysql database object.
func (p *Mysql) InitDB(connectionString string) error {
	db, err := InitDatabase("mysql", connectionString)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// DB get database object.
func (p *Mysql) DB() *sql.DB {
	return p.db
}

// TableNamesQuery returns query to list table names in the current schema
// (not including system tables).
func (p *Mysql) TableNamesQuery() string {
	return `SELECT table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' AND
		table_schema = database()`
}

// ViewNamesQuery returns a list of all view names in the current schema
// (not including system views).
func (p *Mysql) ViewNamesQuery() string {
	return `SELECT table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema = database()`
}

// ObjectQuery returns details query table/veiw of one object.
func (p *Mysql) ObjectQuery() string {
	return `SELECT * FROM %s LIMIT 0`
}
