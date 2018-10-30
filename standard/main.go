package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/tarekbadrshalaan/modelgen/standard/api"
	"github.com/tarekbadrshalaan/modelgen/standard/db"
)

func main() {
	fmt.Println("Go ORM Tutorial")

	if err := db.InitDB("postgres", "host=127.0.0.1 port=5432 user=tarek password=123 dbname=dvdrental sslmode=disable"); err != nil {
		panic(err)
	}
	r := api.ConfigRouter()
	log.Fatal(http.ListenAndServe(":8899", r))

	db.Close()
}
