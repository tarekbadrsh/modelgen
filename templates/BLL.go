package templates

// bllTmpl : template of BLL
var bllTmpl = `package bll

import (
	"strconv"

	"{{.PackageName}}/dal"
	"{{.PackageName}}/dto"
)

{{ if (ne .IDType "string")}}
// Convert{{.IDName}} : covnert {{.IDName}} string to {{.IDName}} {{.IDType}}.
func Convert{{.IDName}}(str string) ({{.IDType}}, error) {
	pram, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	id := {{.IDType}}(pram)
	return id, nil
}
{{ end }}

// GetAll{{pluralize .StructName}} : get All {{pluralizeLower .StructName}}.
func GetAll{{pluralize .StructName}}() ([]*dto.{{DTO .StructName}}, error) {
	{{pluralizeLower .StructName}} := dal.GetAll{{pluralize .StructName}}()
	return dto.{{.StructName}}DALToDTOArr({{pluralizeLower .StructName}})
}

// Get{{.StructName}} : get one {{.TableName}} by id.
func Get{{.StructName}}(id {{.IDType}}) (*dto.{{DTO .StructName}}, error) {
	{{.ShortStructName}}, err := dal.Get{{.StructName}}(id)
	if err != nil {
		return nil, err
	}
	return dto.{{.StructName}}DALToDTO({{.ShortStructName}})
}

{{ if .IsTable}}
// Create{{.StructName}} : create new {{.TableName}}.
func Create{{.StructName}}({{.ShortStructName}} *dto.{{DTO .StructName}}) (*dto.{{DTO .StructName}}, error) {
	{{.TableName}}, err := {{.ShortStructName}}.{{.StructName}}DTOToDAL()
	if err != nil {
		return nil, err
	}
	new{{.TableName}}, err := dal.Create{{.StructName}}({{.TableName}})
	if err != nil {
		return nil, err
	}
	return dto.{{.StructName}}DALToDTO(new{{.TableName}})
}

// Update{{.StructName}} : update exist {{.TableName}}.
func Update{{.StructName}}({{.ShortStructName}} *dto.{{DTO .StructName}}) (*dto.{{DTO .StructName}}, error) {
	{{.TableName}}, err := {{.ShortStructName}}.{{.StructName}}DTOToDAL()
	if err != nil {
		return nil, err
	}
	update{{.TableName}}, err := dal.Update{{.StructName}}({{.TableName}})
	if err != nil {
		return nil, err
	}
	return dto.{{.StructName}}DALToDTO(update{{.TableName}})
}

// Delete{{.StructName}} : delete {{.TableName}} by id.
func Delete{{.StructName}}(id {{.IDType}}) error {
	return dal.Delete{{.StructName}}(id)
}
{{ end }}

`
