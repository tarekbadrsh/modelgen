package database

import (
	"database/sql"

	// Sqlite driver
	_ "github.com/mattn/go-sqlite3"
)

// Sqlite driver.
type Sqlite struct {
	db *sql.DB
}

// GoImport : return import path of go package.
func (p *Sqlite) GoImport() string {
	return "_ \"github.com/mattn/go-sqlite3\""
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

// PrimarykeyQuery interface of Get PrimaryKey Column Name.
func (p *Sqlite) PrimarykeyQuery() string {
	return `WITH RECURSIVE split(content, last, rest) AS (
		VALUES('', '', (select substr(q,0,instr(q, ')')) as pr from (
		  select substr(sql, instr(sql, 'PRIMARY KEY (')+13) as q from sqlite_master where type='table' and name='languages'
		)))
		UNION ALL
		  SELECT
		
			CASE WHEN last = ','
					THEN
						substr(rest, 1, 1)
					ELSE
						content || substr(rest, 1, 1)
			END,
			 substr(rest, 1, 1),
			 substr(rest, 2)
		  FROM split
		  WHERE rest <> ''
		)
		SELECT
			  ltrim(REPLACE(content,',','')) AS 'ValueSplit'
		FROM
			   split
		WHERE
			   last = ',' OR rest ='';
	`
}
