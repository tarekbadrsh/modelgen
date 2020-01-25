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
	AppName            string `json:"App_Name"`
	TargetDirectory    string `json:"Target_Directory"`
	Module             string `json:"Module"`
	DBConnectionString string `json:"DB_CONNECTION_STRING"`
	DBEngine           string `json:"DB_ENGINE"`
	WebAddress         string `json:"API_ADDRESS"`
	WebPort            int    `json:"API_PORT"`
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
		{dir: c.TargetDirectory + "/dal", filepath: "%v/dal/%vDAL.go", tmpfunc: templates.DALTemplate},
		{dir: c.TargetDirectory + "/bll", filepath: "%v/bll/%vBLL.go", tmpfunc: templates.BLLTemplate},
		{dir: c.TargetDirectory + "/dto", filepath: "%v/dto/%vDTO.go", tmpfunc: templates.DTOTemplate},
		{dir: c.TargetDirectory + "/api", filepath: "%v/api/%vAPI.go", tmpfunc: templates.APITemplate},
		{dir: c.TargetDirectory + "/test", filepath: "%v/test/%v_test.go", tmpfunc: templates.TestTemplate, dbImport: c.DBImport},
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
				err = generateFile(m.dir, fmt.Sprintf(m.filepath, c.TargetDirectory, st.StructName), tm, st)
				if err != nil {
					panic(err)
				}
				apiRouters[st.StructName] = true
			}
		}
	}

	singlefile := []gen{
		{dir: c.TargetDirectory + "/db", filepath: c.TargetDirectory + "/db/database.go", tmpfunc: templates.DBTemplate},
		{dir: c.TargetDirectory + "/test", filepath: c.TargetDirectory + "/test/test.json", tmpfunc: templates.ConfigjsonTemplate, data: c},
		{dir: c.TargetDirectory + "/test", filepath: c.TargetDirectory + "/test/config_test.go", tmpfunc: templates.TestConfigTemplate},
		{dir: c.TargetDirectory + "/config", filepath: c.TargetDirectory + "/config/config.go", tmpfunc: templates.ConfigTemplate},

		{dir: c.TargetDirectory + "/logger", filepath: c.TargetDirectory + "/logger/logger.go", tmpfunc: templates.IloggerTemplate},
		{dir: c.TargetDirectory + "/logger", filepath: c.TargetDirectory + "/logger/empty.go", tmpfunc: templates.EmptyloggerTemplate},
		{dir: c.TargetDirectory + "/logger", filepath: c.TargetDirectory + "/logger/zap.go", tmpfunc: templates.ZaploggerTemplate},

		{filepath: c.TargetDirectory + "/api/router.go", tmpfunc: templates.APIRouterTemplate, data: apiRouters},
		{filepath: c.TargetDirectory + "/config.json", tmpfunc: templates.ConfigjsonTemplate, data: c},
		{filepath: c.TargetDirectory + "/main.go", tmpfunc: templates.MainTemplate, data: c},
		{filepath: c.TargetDirectory + "/go.mod", tmpfunc: templates.ModuleTemplate, data: c},
		{filepath: c.TargetDirectory + "/Dockerfile", tmpfunc: templates.DockerTemplate, data: c},
		{filepath: c.TargetDirectory + "/.env", tmpfunc: templates.EnvTemplate, data: c},
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
