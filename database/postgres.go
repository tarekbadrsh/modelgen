package database

import (
	"database/sql"
)

// Postgres driver.
type Postgres struct {
	db *sql.DB
}

// InitDB : initialize postgres database object.
func (p *Postgres) InitDB(connectionString string) error {
	db, err := InitDatabase("postgres", connectionString)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// DB get database object.
func (p *Postgres) DB() *sql.DB {
	return p.db
}

// TableNamesQuery returns query to list table names in the current schema
// (not including system tables).
func (p *Postgres) TableNamesQuery() string {
	return `SELECT table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' AND
		table_schema = current_schema()`
}

// ViewNamesQuery returns a list of all view names in the current schema
// (not including system views).
func (p *Postgres) ViewNamesQuery() string {
	return `SELECT table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema = current_schema()`
}

// ObjectQuery returns details query table/veiw of one object.
func (p *Postgres) ObjectQuery() string {
	return `SELECT * FROM %s LIMIT 0`
}

// PrimarykeyQuery interface of Get PrimaryKey Column Name.
func (p *Postgres) PrimarykeyQuery() string {
	return `SELECT c.column_name
	FROM information_schema.key_column_usage AS c
	LEFT JOIN information_schema.table_constraints AS t
	ON t.constraint_name = c.constraint_name
	WHERE t.table_name = '%s' AND t.constraint_type = 'PRIMARY KEY';`
}
