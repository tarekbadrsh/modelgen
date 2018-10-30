// Package database provides access to database schema metadata, for database/sql drivers.
//
//ref. https://github.com/jimsmart/schema
package database

import (
	"database/sql"
	"fmt"
)

// names queries the database schema metadata and returns
// either a list of table or view names.
//
// It uses the database driver name and the passed query type
// to lookup the appropriate dialect and query.
func names(db *sql.DB, query string) ([]string, error) {
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Scan result into list of names.
	var names []string
	n := ""
	for rows.Next() {
		err = rows.Scan(&n)
		if err != nil {
			return nil, err
		}
		names = append(names, n)
	}
	return names, nil
}

// TableNames returns a list of all table names in the current schema
// (not including system tables).
func TableNames(db TableNamesQuery) ([]string, error) {
	return names(db.DB(), db.TableNamesQuery())
}

// Table returns the column type metadata for the given table name.
func Table(db TableNamesQuery, name string) ([]*sql.ColumnType, error) {
	return object(db, name)
}

// Tables returns column type metadata for all tables in the current schema
// (not including system tables). The returned map is keyed by table name.
//func Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
func Tables(db TableNamesQuery) (map[string][]*sql.ColumnType, error) {
	tablesNames, err := TableNames(db)
	if err != nil {
		return nil, err
	}
	return objects(db, tablesNames)
}

// ViewNames returns a list of all view names in the current schema
// (not including system views).
func ViewNames(db ViewNamesQuery) ([]string, error) {
	return names(db.DB(), db.ViewNamesQuery())
}

// View returns the column type metadata for the given view name.
func View(db ViewNamesQuery, name string) ([]*sql.ColumnType, error) {
	return object(db, name)
}

// Views returns column type metadata for all views in the current schema
// (not including system views). The returned map is keyed by view name.
func Views(db ViewNamesQuery) (map[string][]*sql.ColumnType, error) {
	viewNames, err := ViewNames(db)
	if err != nil {
		return nil, err
	}
	return objects(db, viewNames)
}

// Object queries the database and returns column type metadata
// for a single table or view.
//
// It uses the database driver name to look up the appropriate
// dialect, and the passed table/view name to build the query.
func object(db DBQuery, name string) ([]*sql.ColumnType, error) {
	query := fmt.Sprintf(db.ObjectQuery(), name)
	rows, err := db.DB().Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return rows.ColumnTypes()
}

// objects queries the database and returns metadata about the
// column types for all tables or all views.
//
// It uses the passed list provider function to obtain table/view names,
// and calls object() to fetch the column metadata for each name in the list.
func objects(db DBQuery, objnames []string) (map[string][]*sql.ColumnType, error) {
	if len(objnames) == 0 {
		return nil, nil
	}
	m := make(map[string][]*sql.ColumnType, len(objnames))
	for _, n := range objnames {
		ci, err := object(db, n)
		if err != nil {
			return nil, err
		}
		m[n] = ci
	}
	return m, nil
}

// Primarykeys queries the database and returns list PrimaryKey Columns.
func Primarykeys(db PrimarykeyQuery, name string) ([]string, error) {
	query := fmt.Sprintf(db.PrimarykeyQuery(), name)
	return names(db.DB(), query)
}
