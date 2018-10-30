package modelgen

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tarekbadrshalaan/modelgen/dbutils"
)

// ModelInfo Object pass to Generator, to Generate DTO,DAL.
type ModelInfo struct {
	PackageName     string
	StructNameDAL   string
	StructNameDTO   string
	ShortStructName string
	IsTable         bool
	TableName       string
	IsView          bool
	ViewName        string
	Fields          []Field
	Import          map[string]bool
}

// GenerateStruct generates a struct for the given table.
func GenerateStruct(pkgName, tableName string, viewName string, cols []*sql.ColumnType, primarykeys []string) *ModelInfo {

	structName := dbutils.FmtFieldName(tableName)
	var modelInfo = &ModelInfo{
		PackageName:     pkgName,
		StructNameDAL:   fmt.Sprintf("%vDAL", structName),
		StructNameDTO:   fmt.Sprintf("%vDTO", structName),
		ShortStructName: strings.ToLower(string(structName[0])),
		TableName:       tableName,
		ViewName:        viewName,
	}
	modelInfo.IsTable = tableName != ""
	modelInfo.IsView = viewName != ""

	modelInfo.Import = make(map[string]bool)
	for _, col := range cols {
		field := colToField(col, primarykeys)
		modelInfo.Fields = append(modelInfo.Fields, field)

		if strings.HasPrefix(field.GoType.String(), "sql") {
			modelInfo.Import["database/sql"] = true
		}
		if strings.HasPrefix(field.GoType.String(), "time") {
			modelInfo.Import["time"] = true
		}
	}
	return modelInfo
}
