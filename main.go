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
		{dir: "Application/dal", filepath: "Application/dal/%vDAL.go", tmpfunc: templates.DALTemplate},
		{dir: "Application/bll", filepath: "Application/bll/%vBLL.go", tmpfunc: templates.BLLTemplate},
		{dir: "Application/dto", filepath: "Application/dto/%vDTO.go", tmpfunc: templates.DTOTemplate},
		{dir: "Application/api", filepath: "Application/api/%vAPI.go", tmpfunc: templates.APITemplate},
		{dir: "Application/test", filepath: "Application/test/%v_test.go", tmpfunc: templates.TestTemplate, dbImport: c.DBImport},
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
				err = generateFile(m.dir, fmt.Sprintf(m.filepath, st.StructName), tm, st)
				if err != nil {
					panic(err)
				}
				apiRouters[st.StructName] = true
			}
		}
	}

	singlefile := []gen{
		{dir: "Application/db", filepath: "Application/db/database.go", tmpfunc: templates.DBTemplate},
		{dir: "", filepath: "Application/api/router.go", tmpfunc: templates.APIRouterTemplate, data: apiRouters},
		{dir: "", filepath: "Application/config.json", tmpfunc: templates.ConfigTemplate, data: c},
		{dir: "Application/test", filepath: "Application/test/test.json", tmpfunc: templates.ConfigTemplate, data: c},
		{dir: "Application/test", filepath: "Application/test/config_test.go", tmpfunc: templates.TestConfigTemplate},
		{dir: "", filepath: "Application/main.go", tmpfunc: templates.MainTemplate, data: c},
		{dir: "", filepath: "Application/go.mod", tmpfunc: templates.ModuleTemplate, data: c},
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
	if dir != "" { // if the path already exist, don't neet to create again.
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
