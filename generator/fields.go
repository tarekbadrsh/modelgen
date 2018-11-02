package generator

import (
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	"github.com/tarekbadrshalaan/modelgen/dbutils"
)

// Field collection of var about one column.
type Field struct {
	DatabaseName string
	SQLType      string
	IsPrimaryKey bool
	GoType       reflect.Type
	GoName       string
	DALfmt       string
	DTOfmt       string
}

// colToField  : Convert Column to Field
func colToField(col *sql.ColumnType, primarykeys []string) Field {

	f := Field{
		DatabaseName: col.Name(),
		SQLType:      col.DatabaseTypeName(),
		GoType:       col.ScanType(),
		GoName:       dbutils.FmtFieldName(col.Name()),
	}
	for _, p := range primarykeys {
		if p == f.DatabaseName {
			f.IsPrimaryKey = true
			break
		}
	}
	var annotations []string
	annotations = append(annotations, fmt.Sprintf("json:\"%s\"", f.DatabaseName))
	f.DTOfmt = fmt.Sprintf("%s %s `%s`", f.GoName, f.GoType, strings.Join(annotations, " "))

	gormannotations := fmt.Sprintf("gorm:\"column:%s\"", f.DatabaseName)
	if f.IsPrimaryKey {
		gormannotations = fmt.Sprintf("gorm:\"column:%s;primary_key:true\"", f.DatabaseName)
	}
	annotations = append(annotations, gormannotations)
	f.DALfmt = fmt.Sprintf("%s %s `%s`", f.GoName, f.GoType, strings.Join(annotations, " "))

	return f
}
