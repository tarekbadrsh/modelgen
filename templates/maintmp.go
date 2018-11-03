package templates

// mainTmpl : template of mainTmpl
var mainTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"

	{{.DBImport}}
	"github.com/tarekbadrshalaan/goStuff/configuration"
	"{{.Module}}/api"
	"{{.Module}}/db"
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

	// database.
	if err := db.InitDB(c.DBEngine, c.DBConnectionString); err != nil {
		panic(err)
	}
	defer db.Close()
	// database.

	// webserver.
	r := api.ConfigRouter()
	addr := fmt.Sprintf("%v:%d", c.WebAddress, c.WebPort)
	log.Fatal(http.ListenAndServe(addr, r))
	// webserver.
}

`
