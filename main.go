package main

import (
	"fmt"
	"os"

	"text/template"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
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
}

type gen struct {
	dir      string
	pkgname  string
	filepath string
	tmpfunc  func() (*template.Template, error)
	data     interface{}
}

func main() {
	// configruation.
	c := &config{}
	if err := configuration.JSON("config.json", c); err != nil {
		panic(err)
	}
	// configruation.

	// database.
	p := database.GetDatabaseEngine(c.DBEngine)
	if err := p.InitDB(c.DBConnectionString); err != nil {
		panic(err)
	}
	// database.

	mutifiles := []gen{
		{dir: "Application/dal", pkgname: "dal", filepath: "Application/dal/%vDAL.go", tmpfunc: templates.DALTemplate},
		{dir: "Application/bll", pkgname: "bll", filepath: "Application/bll/%vBLL.go", tmpfunc: templates.BLLTemplate},
		{dir: "Application/dto", pkgname: "dto", filepath: "Application/dto/%vDTO.go", tmpfunc: templates.DTOTemplate},
		{dir: "Application/api", pkgname: "api", filepath: "Application/api/%vAPI.go", tmpfunc: templates.APITemplate},
	}

	tables, err := database.Tables(p)
	if err != nil {
		panic(err)
	}
	apiRouters := map[string]bool{}
	for _, m := range mutifiles {
		for table, cols := range tables {
			primarykeys, err := database.Primarykeys(p, table)
			if err != nil {
				panic(err)
			}
			if len(primarykeys) == 1 {
				st := generator.GenerateStruct(c.Module, table, "", cols, primarykeys)
				err = generateFile(m.dir, fmt.Sprintf(m.filepath, st.StructName), m.tmpfunc, st)
				if err != nil {
					panic(err)
				}
				apiRouters[st.StructName] = true
			}
		}
	}

	singlefile := []gen{
		{dir: "Application/db", filepath: "Application/db/database.go", tmpfunc: templates.DBTemplate, data: nil},
		{dir: "", filepath: "Application/api/router.go", tmpfunc: templates.APIRouterTemplate, data: apiRouters},
		{dir: "", filepath: "Application/config.json", tmpfunc: templates.ConfigTemplate, data: c},
		{dir: "", filepath: "Application/main.go", tmpfunc: templates.MainTemplate, data: c.Module},
		{dir: "", filepath: "Application/go.mod", tmpfunc: templates.ModuleTemplate, data: c.Module},
	}
	for _, s := range singlefile {
		err = generateFile(s.dir, s.filepath, s.tmpfunc, s.data)
		if err != nil {
			panic(err)
		}
	}
}

func generateFile(dir, filepath string, tmpfunc func() (*template.Template, error), data interface{}) error {
	if dir != "" { // if the path already exist, don't neet to create again.
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {
			return err
		}
	}
	tm, err := tmpfunc()
	if err != nil {
		return err
	}
	return generateTemplate(data, filepath, tm)
}

func generateTemplate(data interface{}, filePath string, tmpl *template.Template) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	if err = tmpl.Execute(f, data); err != nil {
		return err
	}
	return nil
}
