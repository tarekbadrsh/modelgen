package database

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

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

// DBDriver interface of database object and Driver.
type DBDriver interface {
	DB() *sql.DB
	InitDB(connectionString string) (*sql.DB, error)
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
