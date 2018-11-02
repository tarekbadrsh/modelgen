package templates

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

// DALTemplate : return template of DAL.
func DALTemplate() (*template.Template, error) {
	return getTemplate(DALTmpl)
}

func getTemplate(temp string) (*template.Template, error) {
	var funcMap = template.FuncMap{
		"pluralize":      inflection.Plural,
		"pluralizeLower": pluralizeLower,
		"title":          strings.Title,
		"toLower":        strings.ToLower,
		"DTO":            structNameDTO,
		"DAL":            structNameDAL,
	}

	tmpl, err := template.New("model").Funcs(funcMap).Parse(temp)

	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func pluralizeLower(s string) string {
	return strings.ToLower(inflection.Plural(s))
}

func structNameDTO(s string) string {
	return fmt.Sprintf("%vDTO", s)
}

func structNameDAL(s string) string {
	return fmt.Sprintf("%vDAL", s)
}
