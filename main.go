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
	DBConnectionString string
	DBEngine           string
	WebAddress         string
	WebPort            int
}

func main() {
	// configruation.
	c := &config{}
	err := configuration.JSON("config.json", c)
	if err != nil {
		panic(err)
	}
	// configruation.
	pkgPath := "github.com/tarekbadrshalaan/modelgen/Application"
	p := &database.Postgres{}
	err = p.InitDB(c.DBConnectionString)
	if err != nil {
		panic(err)
	}
	apiRouters := []string{}

	mod := []struct {
		dir            string
		pkgname        string
		filename       string
		generateTables bool
		data           interface{}
		tmpfunc        func() (*template.Template, error)
	}{
		{dir: "Application/dal", generateTables: true, pkgname: "dal", filename: "DAL.go", tmpfunc: templates.DALTemplate, data: nil},
		{dir: "Application/bll", generateTables: true, pkgname: "bll", filename: "BLL.go", tmpfunc: templates.BLLTemplate, data: nil},
		{dir: "Application/dto", generateTables: true, pkgname: "dto", filename: "DTO.go", tmpfunc: templates.DTOTemplate, data: nil},
		{dir: "Application/api", generateTables: true, pkgname: "api", filename: "API.go", tmpfunc: templates.APITemplate, data: nil},
		{dir: "Application/db", generateTables: false, pkgname: "db", filename: "db.go", tmpfunc: templates.DBTemplate, data: nil},
		{dir: "Application/db", generateTables: false, pkgname: "api", filename: "router.go", tmpfunc: templates.APIRouterTemplate, data: apiRouters},
	}
	for _, m := range mod {
		err := os.MkdirAll(m.dir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

	tables, err := database.Tables(p)
	if err != nil {
		panic(err)
	}

	for table, cols := range tables {
		primarykeys, err := database.Primarykeys(p, table)
		if err != nil {
			panic(err)
		}
		st := generator.GenerateStruct(pkgPath, table, "", cols, primarykeys)
		if len(primarykeys) == 1 {
			for _, m := range mod {
				if !m.generateTables {
					continue
				}
				tm, err := m.tmpfunc()
				if err != nil {
					panic(err)
				}
				err = generatefile(st, fmt.Sprintf("Application/%v/%v%v", m.pkgname, st.StructName, m.filename), tm)
				if err != nil {
					panic(err)
				}
			}
			apiRouters = append(apiRouters, st.StructName)
		}
	}
	router, err := templates.APIRouterTemplate()
	if err != nil {
		panic(err)
	}
	err = generatefile(apiRouters, "Application/api/router.go", router)
	if err != nil {
		panic(err)
	}

	dbtmp, err := templates.DBTemplate()
	if err != nil {
		panic(err)
	}
	err = generatefile(nil, "Application/db/database.go", dbtmp)
	if err != nil {
		panic(err)
	}

	conf, err := templates.ConfigTemplate()
	if err != nil {
		panic(err)
	}
	err = generatefile(c, "Application/config.json", conf)
	if err != nil {
		panic(err)
	}

	maintmp, err := templates.MainTemplate()
	if err != nil {
		panic(err)
	}
	err = generatefile(pkgPath, "Application/main.go", maintmp)
	if err != nil {
		panic(err)
	}

}

func generatefile(data interface{}, filePath string, tmpl *template.Template) error {
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
