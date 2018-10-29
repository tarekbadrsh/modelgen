package templates

import (
	"text/template"
)

// DALTemplate : return template of DAL.
func DALTemplate() (*template.Template, error) {
	return getTemplate(DALTmpl)
}

func getTemplate(temp string) (*template.Template, error) {
	// var funcMap = template.FuncMap{
	// 	"pluralize":        inflection.Plural,
	// 	"title":            strings.Title,
	// 	"toLower":          strings.ToLower,
	// 	"toLowerCamelCase": camelToLowerCamel,
	// 	"toSnakeCase":      snaker.CamelToSnake,
	// }

	//tmpl, err := template.New("model").Funcs(funcMap).Parse(t)
	tmpl, err := template.New("model").Parse(temp)

	if err != nil {
		return nil, err
	}
	return tmpl, nil
}
