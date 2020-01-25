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
	"{{.Module}}/logger"
)

func main() {
	/* configuration initialize start */
	c, err := config.GetConfigs()
	if err != nil {
		panic(err)
	}
	/* configuration initialize end */

	/* logger initialize start */
	mylogger := logger.NewZapLogger()
	logger.InitLogger(&mylogger)
	defer logger.Close()
	/* logger initialize end */

	/* initialize database start */
	if err := db.InitDB(c.DBEngine, c.DBConnectionString); err != nil {
		panic(err)
	}
	defer db.Close()
	/* initialize database end */

	// webserver.
	r := api.ConfigRouter()
	addr := fmt.Sprintf("%v:%d", c.WebAddress, c.WebPort)
	log.Fatal(http.ListenAndServe(addr, r))
	// webserver.
}

`
