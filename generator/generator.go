package generator

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/tarekbadrshalaan/modelgen/dbutils"
)

// ModelInfo Object pass to Generator, to Generate DTO,DAL.
type ModelInfo struct {
	Module          string
	StructName      string
	ShortStructName string
	IsTable         bool
	TableName       string
	IsView          bool
	ViewName        string
	Fields          []Field
	Import          map[string]bool
	IDName          string
	IDType          string
	DBImport        string
	JSONobj         string
}

// GenerateStruct generates a struct for the given table.
func GenerateStruct(module, tableName string, viewName string, cols []*sql.ColumnType, primarykeys []string, dbImport string) *ModelInfo {

	structName := dbutils.FmtFieldName(tableName)
	var modelInfo = &ModelInfo{
		Module:          module,
		StructName:      structName,
		ShortStructName: strings.ToLower(string(structName[0])),
		TableName:       tableName,
		ViewName:        viewName,
		DBImport:        dbImport,
	}
	modelInfo.IsTable = tableName != ""
	modelInfo.IsView = viewName != ""
	modelInfo.JSONobj = "{"
	modelInfo.Import = make(map[string]bool)
	for i, col := range cols {
		field := colToField(col, primarykeys)
		modelInfo.Fields = append(modelInfo.Fields, field)

		if field.IsPrimaryKey {
			modelInfo.IDName = field.GoName
			modelInfo.IDType = field.GoType.String()
		}
		js := fmt.Sprintf("\"%v\":\"\",", field.DatabaseName)
		if i == len(cols)-1 {
			js = fmt.Sprintf("\"%v\":\"\"", field.DatabaseName)
		}
		modelInfo.JSONobj += js

		if strings.HasPrefix(field.GoType.String(), "sql") {
			modelInfo.Import["database/sql"] = true
		}
		if strings.HasPrefix(field.GoType.String(), "time") {
			modelInfo.Import["time"] = true
		}
	}
	modelInfo.JSONobj += "}"
	return modelInfo
}
