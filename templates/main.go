package templates

// mainTmpl : template of mainTmpl
var mainTmpl = `package main

import (
	"fmt"
	"log"
	"net/http"

	{{.DBImport}}
	"{{.Module}}/api"
	"{{.Module}}/config"
	"{{.Module}}/db"
)

func main() {
	/* configuration initialize start */
	c := config.GetConfigs()
	/* configuration initialize end */

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
