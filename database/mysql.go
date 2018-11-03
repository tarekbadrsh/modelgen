package database

import (
	"database/sql"

	// Mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Mysql driver.
type Mysql struct {
	db *sql.DB
}

// GoImport : return import path of go package.
func (p *Mysql) GoImport() string {
	return "_ \"github.com/go-sql-driver/mysql\""
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

// PrimarykeyQuery interface of Get PrimaryKey Column Name.
func (p *Mysql) PrimarykeyQuery() string {
	return `SELECT k.column_name
	FROM information_schema.table_constraints t
	JOIN information_schema.key_column_usage k
	USING(constraint_name,table_name)
	WHERE t.constraint_type='PRIMARY KEY' 
	  AND t.table_name='%s';`
}
