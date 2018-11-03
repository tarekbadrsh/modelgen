package database

import (
	"database/sql"

	// Mssql driver
	_ "github.com/denisenkom/go-mssqldb"
)

// Mssql driver.
type Mssql struct {
	db *sql.DB
}

// GoImport : return import path of go package.
func (p *Mssql) GoImport() string {
	return "_ \"github.com/denisenkom/go-mssqldb\""
}

// InitDB : initialize Mssql database object.
func (p *Mssql) InitDB(connectionString string) error {
	db, err := InitDatabase("mssql.MssqlDriver", connectionString)
	if err != nil {
		return err
	}
	p.db = db
	return nil
}

// DB get database object.
func (p *Mssql) DB() *sql.DB {
	return p.db
}

// TableNamesQuery returns query to list table names in the current schema
// (not including system tables).
func (p *Mssql) TableNamesQuery() string {
	return `SELECT T.name as name
	FROM
		sys.tables AS T
		INNER JOIN sys.schemas AS S ON S.schema_id = T.schema_id
		LEFT JOIN sys.extended_properties AS EP ON EP.major_id = T.[object_id]
	WHERE
		T.is_ms_shipped = 0 AND
		(EP.class_desc IS NULL OR (EP.class_desc <> 'OBJECT_OR_COLUMN' AND
		EP.[name] <> 'microsoft_database_tools_support'))`
}

// ViewNamesQuery returns a list of all view names in the current schema
// (not including system views).
func (p *Mssql) ViewNamesQuery() string {
	return `SELECT T.name as name
	FROM
		sys.views AS T
		INNER JOIN sys.schemas AS S ON S.schema_id = T.schema_id
		LEFT JOIN sys.extended_properties AS EP ON EP.major_id = T.[object_id]
	WHERE
		T.is_ms_shipped = 0 AND
		(EP.class_desc IS NULL OR (EP.class_desc <> 'OBJECT_OR_COLUMN' AND
		EP.[name] <> 'microsoft_database_tools_support'))`
}

// ObjectQuery returns details query table/veiw of one object.
func (p *Mssql) ObjectQuery() string {
	return `SELECT * FROM %s WHERE 1=0`
}

// PrimarykeyQuery interface of Get PrimaryKey Column Name.
func (p *Mssql) PrimarykeyQuery() string {
	return `SELECT COLUMN_NAME
	FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE
	WHERE OBJECTPROPERTY(OBJECT_ID(CONSTRAINT_SCHEMA + '.' + QUOTENAME(CONSTRAINT_NAME)), 'IsPrimaryKey') = 1
	AND TABLE_NAME = '%s' AND TABLE_SCHEMA = 'Schema'
	`
}
