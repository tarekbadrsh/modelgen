package templates

// DALTmpl : template of DAL
var DALTmpl = `package dal

import ( 
	{{ range $key, $value := .Import}}"{{$key}}"
	{{ end }}
	"{{.PackageName}}/db"
) 

// {{DAL .StructName}} : data access layer ({{.TableName}}) table.
type {{DAL .StructName}} struct {
	{{range .Fields}}{{.DALfmt}}
	{{end}}
}

// TableName sets the insert table name for this struct type
func ({{.ShortStructName}} *{{DAL .StructName}}) TableName() string {
	return "{{.TableName}}"
} 

// GetAll{{pluralize .StructName}} : get all {{pluralizeLower .StructName}}.
func GetAll{{pluralize .StructName}}() []*{{DAL .StructName}} {
	{{pluralizeLower .StructName}} := []*{{DAL .StructName}}{}
	db.DB().Find(&{{pluralizeLower .StructName}})
	return {{pluralizeLower .StructName}}
}

// Get{{.StructName}} : get one {{.TableName}} by id.
func Get{{.StructName}}(id {{.IDType}}) (*{{DAL .StructName}}, error) {
	{{.ShortStructName}} := &{{DAL .StructName}}{}
	result := db.DB().First({{.ShortStructName}}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return {{.ShortStructName}}, nil
}

// Create{{.StructName}} : create new {{.TableName}}.
func Create{{.StructName}}({{.ShortStructName}} *{{DAL .StructName}}) (*{{DAL .StructName}}, error) {
	result := db.DB().Create({{.ShortStructName}})
	if result.Error != nil {
		return nil, result.Error
	}
	return {{.ShortStructName}}, nil
}

// Update{{.StructName}} : update exist {{.TableName}}.
func Update{{.StructName}}({{.ShortStructName}} *{{DAL .StructName}}) (*{{DAL .StructName}}, error) {
	_, err := Get{{.StructName}}({{.ShortStructName}}.{{.IDName}})
	if err != nil {
		return nil, err
	}
	result := db.DB().Save({{.ShortStructName}})
	if result.Error != nil {
		return nil, result.Error
	}
	return {{.ShortStructName}}, nil
}

// Delete{{.StructName}} : delete {{.TableName}} by id.
func Delete{{.StructName}}(id {{.IDType}}) error {
	{{.ShortStructName}}, err := Get{{.StructName}}(id)
	if err != nil {
		return err
	}
	result := db.DB().Delete({{.ShortStructName}})
	return result.Error
}

`
