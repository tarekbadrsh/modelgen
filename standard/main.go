package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/modelgen/standard/api"
	"github.com/tarekbadrshalaan/modelgen/standard/db"
)

func main() {
	// configruation
	if err := db.InitDB("postgres", "host=127.0.0.1 port=5432 user=tarek password=123 dbname=dvdrental sslmode=disable"); err != nil {
		panic(err)
	}
	defer db.Close()
	r := api.ConfigRouter()
	log.Fatal(http.ListenAndServe(":8899", r))
}
