package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/goStuff/configuration"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/api"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/db"
)

type config struct {
	DBConnectionString string
	DBEngine           string
	WebAddress         string
	WebPort            int
}

func main() {
	// configurations.
	c := &config{}
	err := configuration.JSON("config.json", c)
	if err != nil {
		panic(err)
	}
	// configurations.

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

