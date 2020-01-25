package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/api"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/config"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/db"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
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

