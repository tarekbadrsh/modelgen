package templates

// DALTmpl : template of DAL
var DALTmpl = `package {{.PackageName}}

import ( 
	{{ range $key, $value := .Import}}"{{$key}}"
	{{ end }}
) 

// {{.StructNameDAL}} : 
type {{.StructNameDAL}} struct {
	{{range .Fields}}{{.DALfmt}}
    {{end}}
}

// TableName sets the insert table name for this struct type
func ({{.ShortStructName}} *{{.StructNameDAL}}) TableName() string {
	return "{{.TableName}}"
}
`
