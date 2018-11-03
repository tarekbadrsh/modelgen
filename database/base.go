package database

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/pkg/errors"
)

// DBDriver interface of database object and Driver.
type DBDriver interface {
	// DB get database object.
	DB() *sql.DB

	// InitDB : initialize database object.
	InitDB(connectionString string) error

	// GoImport : return import path of go package.
	GoImport() string

	// returns details query table/veiw of one object.
	ObjectQuery() string

	// returns query to list table names in the current schema.
	TableNamesQuery() string

	// returns a list of all view names in the current schema.
	ViewNamesQuery() string

	// return query Get PrimaryKey Column Name.
	PrimarykeyQuery() string
}

// DBQuery :.
type DBQuery interface {
	DB() *sql.DB
	ObjectQuery() string
}

// TableNamesQuery interface of DBDriver and TableNamesQuery.
type TableNamesQuery interface {
	DB() *sql.DB
	ObjectQuery() string
	TableNamesQuery() string
}

// ViewNamesQuery interface of DBDriver and ViewNamesQuery.
type ViewNamesQuery interface {
	DB() *sql.DB
	ObjectQuery() string
	ViewNamesQuery() string
}

// PrimarykeyQuery return query Get PrimaryKey Column Name.
type PrimarykeyQuery interface {
	DB() *sql.DB
	PrimarykeyQuery() string
}

// GetDatabaseEngine : return DBEngine By name
func GetDatabaseEngine(DBEngine string) DBDriver {
	switch strings.ToLower(DBEngine) {
	case "postgres":
		return &Postgres{}
	case "mssql":
		return &Mssql{}
	case "mysql":
		return &Mysql{}
	case "oracle":
		return &Oracle{}
	case "sqlite":
		return &Sqlite{}
	}

	return nil
}

// InitDatabase is initialize database object.
func InitDatabase(driverName, connectionString string) (*sql.DB, error) {
	if connectionString == "" {
		return nil, fmt.Errorf("no ConnectionString, please add db ConnectionString")
	}
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, errors.Wrap(err, "Open database")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "Ping database")
	}
	return db, nil
}
