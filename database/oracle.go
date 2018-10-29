package database

import (
	"database/sql"
)

// Oracle driver.
type Oracle struct {
	db *sql.DB
}

// InitDB : initialize Oracle database object.
func (p *Oracle) InitDB(connectionString string) error {
	db, err := InitDatabase("goracle", connectionString)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// DB get database object.
func (p *Oracle) DB() *sql.DB {
	return p.db
}

// TableNamesQuery returns query to list table names in the current schema
// (not including system tables).
func (p *Oracle) TableNamesQuery() string {
	return `SELECT table_name
	FROM
		all_tables
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)`
}

// ViewNamesQuery returns a list of all view names in the current schema
// (not including system views).
func (p *Oracle) ViewNamesQuery() string {
	return `SELECT view_name
	FROM
		all_views
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)`
}

// ObjectQuery returns details query table/veiw of one object.
func (p *Oracle) ObjectQuery() string {
	return `SELECT * FROM %s WHERE 1=0`
}

// PrimarykeyQuery interface of Get PrimaryKey Column Name.
func (p *Oracle) PrimarykeyQuery() string {
	return `SELECT cols.column_name
	FROM all_constraints cons, all_cons_columns cols
	WHERE cols.table_name = '%s'
	AND cons.constraint_type = 'P'
	AND cons.constraint_name = cols.constraint_name`
}
