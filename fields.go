package modelgen

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
	GoType       reflect.Type
	GoName       string
	DALfmt       string
	DTOfmt       string
}

// ColToField  : Convert Column to Field
func ColToField(col *sql.ColumnType) Field {

	f := Field{
		DatabaseName: col.Name(),
		SQLType:      col.DatabaseTypeName(),
		GoType:       col.ScanType(),
		GoName:       dbutils.FmtFieldName(col.Name()),
	}

	var annotations []string
	annotations = append(annotations, fmt.Sprintf("json:\"%s\"", f.DatabaseName))
	annotations = append(annotations, fmt.Sprintf("gorm:\"column:%s\"", f.DatabaseName))
	f.DALfmt = fmt.Sprintf("%s %s `%s`", f.GoName, f.GoType, strings.Join(annotations, " "))
	f.DTOfmt = fmt.Sprintf("%s %s `%s`", f.GoName, f.GoType, strings.Join(annotations, " "))
	return f
}
