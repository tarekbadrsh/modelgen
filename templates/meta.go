package templates

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

// DALTemplate : return template of DAL.
func DALTemplate() (*template.Template, error) {
	return getTemplate(dalTmpl)
}

// DTOTemplate : return template of DTO.
func DTOTemplate() (*template.Template, error) {
	return getTemplate(dtoTmpl)
}

// BLLTemplate : return template of BLL.
func BLLTemplate() (*template.Template, error) {
	return getTemplate(bllTmpl)
}

// APITemplate : return template of API.
func APITemplate() (*template.Template, error) {
	return getTemplate(apiTmpl)
}

// APIRouterTemplate : return template of API.
func APIRouterTemplate() (*template.Template, error) {
	return getTemplate(apiRouterTmpl)
}

// DBTemplate : return template of Database.
func DBTemplate() (*template.Template, error) {
	return getTemplate(dbTmpl)
}

// ConfigTemplate : return template of Config.
func ConfigTemplate() (*template.Template, error) {
	return getTemplate(configTmpl)
}

// MainTemplate : return template of Main.
func MainTemplate() (*template.Template, error) {
	return getTemplate(mainTmpl)
}

// ModuleTemplate : return template of go.mod.
func ModuleTemplate() (*template.Template, error) {
	return getTemplate(moduleTmpl)
}

// TestTemplate : return template of Test.
func TestTemplate() (*template.Template, error) {
	return getTemplate(testTmpl)
}

// TestConfigTemplate : return template of Test Config.
func TestConfigTemplate() (*template.Template, error) {
	return getTemplate(testConfigTmpl)
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
