package templates

import (
	"embed"
	"fmt"
	"strings"
	"text/template"

	"github.com/jinzhu/inflection"
)

//go:embed templates/*
var tmps embed.FS

// APITemplate : return template of API.
func APITemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/api.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// APIRouterTemplate : return template of API.
func APIRouterTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/apirouter.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// BLLTemplate : return template of BLL.
func BLLTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/bll.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// ConfigTemplate : return template of Config.
func ConfigTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/config.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// ConfigjsonTemplate : return template of Config json file.
func ConfigjsonTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/configjson.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// DALTemplate : return template of DAL.
func DALTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/dal.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// DBTemplate : return template of Database.
func DBTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/db.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// DockerTemplate : return template of Dockerfile.
func DockerTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/dockerfile.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// DTOTemplate : return template of DTO.
func DTOTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/dto.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// IloggerTemplate : return template of Ilogger.
func IloggerTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/logger/logger.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// EmptyloggerTemplate : return template of EmptyLogger.
func EmptyloggerTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/logger/empty.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// ZaploggerTemplate : return template of ZapLogger.
func ZaploggerTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/logger/zap.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// EnvTemplate : return template of environment variables file.
func EnvTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/env.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// MainTemplate : return template of Main.
func MainTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/main.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// ModuleTemplate : return template of go.mod.
func ModuleTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/mod.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// TestTemplate : return template of Test.
func TestTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/test.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

// TestConfigTemplate : return template of Test Config.
func TestConfigTemplate() (*template.Template, error) {
	tmpl, err := tmps.ReadFile("templates/testconfig.gotmp")
	if err != nil {
		return nil, err
	}
	return getTemplate(string(tmpl))
}

func getTemplate(temp string) (*template.Template, error) {
	var funcMap = template.FuncMap{
		"pluralize":      inflection.Plural,
		"pluralizeLower": pluralizeLower,
		"title":          strings.Title,
		"toLower":        strings.ToLower,
		"DTO":            structNameDTO,
		"DAL":            structNameDAL,
		"backQuote":      backQuote,
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

func backQuote() string {
	return fmt.Sprint("`")
}
