package main

import (
	"fmt"
	"os"

	"text/template"

	"github.com/tarekbadrshalaan/goStuff/configuration"
	"github.com/tarekbadrshalaan/modelgen/database"
	"github.com/tarekbadrshalaan/modelgen/generator"
	"github.com/tarekbadrshalaan/modelgen/templates"
)

type config struct {
	AppName            string
	Module             string
	DBConnectionString string
	DBEngine           string
	WebAddress         string
	WebPort            int
	DBImport           string
}

type gen struct {
	dir      string
	filepath string
	tmpfunc  func() (*template.Template, error)
	data     interface{}
	dbImport string
}

func main() {
	// configurations.
	c := &config{}
	if err := configuration.JSON("config.json", c); err != nil {
		panic(err)
	}
	// configurations.

	// database.
	p := database.GetDatabaseEngine(c.DBEngine)
	if err := p.InitDB(c.DBConnectionString); err != nil {
		panic(err)
	}
	c.DBImport = p.GoImport()
	// database.

	mutifiles := []gen{
		{dir: c.AppName + "/dal", filepath: "%v/dal/%vDAL.go", tmpfunc: templates.DALTemplate},
		{dir: c.AppName + "/bll", filepath: "%v/bll/%vBLL.go", tmpfunc: templates.BLLTemplate},
		{dir: c.AppName + "/dto", filepath: "%v/dto/%vDTO.go", tmpfunc: templates.DTOTemplate},
		{dir: c.AppName + "/api", filepath: "%v/api/%vAPI.go", tmpfunc: templates.APITemplate},
		{dir: c.AppName + "/test", filepath: "%v/test/%v_test.go", tmpfunc: templates.TestTemplate, dbImport: c.DBImport},
	}

	tables, err := database.Tables(p)
	if err != nil {
		panic(err)
	}

	apiRouters := map[string]bool{}
	for _, m := range mutifiles {
		tm, err := m.tmpfunc()
		if err != nil {
			panic(err)
		}
		for table, cols := range tables {
			primarykeys, err := database.Primarykeys(p, table)
			if err != nil {
				panic(err)
			}
			if len(primarykeys) == 1 {
				st := generator.GenerateStruct(c.Module, table, "", cols, primarykeys, c.DBImport)
				err = generateFile(m.dir, fmt.Sprintf(m.filepath, c.AppName, st.StructName), tm, st)
				if err != nil {
					panic(err)
				}
				apiRouters[st.StructName] = true
			}
		}
	}

	singlefile := []gen{
		{dir: c.AppName + "/db", filepath: c.AppName + "/db/database.go", tmpfunc: templates.DBTemplate},
		{dir: "", filepath: c.AppName + "/api/router.go", tmpfunc: templates.APIRouterTemplate, data: apiRouters},
		{dir: "", filepath: c.AppName + "/config.json", tmpfunc: templates.ConfigTemplate, data: c},
		{dir: c.AppName + "/test", filepath: c.AppName + "/test/test.json", tmpfunc: templates.ConfigTemplate, data: c},
		{dir: c.AppName + "/test", filepath: c.AppName + "/test/config_test.go", tmpfunc: templates.TestConfigTemplate},
		{dir: "", filepath: c.AppName + "/main.go", tmpfunc: templates.MainTemplate, data: c},
		{dir: "", filepath: c.AppName + "/go.mod", tmpfunc: templates.ModuleTemplate, data: c},
	}
	for _, s := range singlefile {
		tm, err := s.tmpfunc()
		if err != nil {
			panic(err)
		}
		err = generateFile(s.dir, s.filepath, tm, s.data)
		if err != nil {
			panic(err)
		}
	}
}

func generateFile(dir, filepath string, tmp *template.Template, data interface{}) error {
	if dir != "" { // if the path already exist, no need to generate again.
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return generateTemplate(data, filepath, tmp)
}

func generateTemplate(data interface{}, filePath string, tmpl *template.Template) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}
