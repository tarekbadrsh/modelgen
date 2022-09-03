# modelgen

[![CircleCI Status](https://img.shields.io/github/release/tarekbadrshalaan/modelgen.svg)](https://github.com/tarekbadrshalaan/modelgen/releases)
[![GoDoc](https://godoc.org/github.com/tarekbadrshalaan/modelgen?status.svg)](https://godoc.org/github.com/tarekbadrshalaan/modelgen)
[![Go Report Card](https://goreportcard.com/badge/github.com/tarekbadrshalaan/modelgen)](https://goreportcard.com/report/github.com/tarekbadrshalaan/modelgen)
[![Build Status](https://travis-ci.org/tarekbadrshalaan/modelgen.svg?branch=master)](https://travis-ci.org/tarekbadrshalaan/modelgen)
[![CircleCI Status](https://circleci.com/gh/tarekbadrshalaan/modelgen.svg?style=shield)](https://circleci.com/gh/tarekbadrshalaan/modelgen)
[![standar-readme compliant](https://img.shields.io/badge/readme%20style-standar-brightgreen.svg)](https://github.com/RichardLitt/standar-readme)




Application to create (start app) go webservice with 3-Tier Architecture.

The Generated Application includes :- 

- Go mod
- Configuration file
- Dockerfile
- DAL,BLL,DTO,API and API_Tests for every Database Table 
- Compatible with `mysql` `postgres` `mssql` `sqlite` `oracle`
- Using [![GORM](https://github.com/jinzhu/gorm)](https://github.com/jinzhu/gorm) as ORM

## Installation

```
$ go get -u github.com/tarekbadrshalaan/modelgen
$ vi config.json
  
  {
    "AppName"               :   "Application",
    "Module"                :   "{{github.com/packagepath}}",
    "DBConnectionString"    :   "{{Database ConnectionString}}",
    "DBEngine"              :   "{{Database Engine}}",
    "WebAddress"            :   "0.0.0.0",
    "WebPort"               :   7070
  }

$ modelgen 
$ cd Application/
$ go build .
  go: finding github.com/tarekbadrshalaan/goStuff/configuration latest
  go: finding github.com/jinzhu/inflection latest
$ ./Application
```

## Example 
- postgres Database
http://www.postgresqltutorial.com/postgresql-sample-database/
to restore the database use : 
- Run postgres on localhost, use this command for restoring database:
```
pg_restore -U postgres -d dvdrental ~/dvdrental.tar
```

- Run postgres into docker container:
```
# run the container
- docker run --rm -it --name pg -p 5454:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_HOST_AUTH_METHOD=trust -e POSTGRES_DB=dvdrental -v $(pwd):/mydata postgres

# open the pg container shell command
- docker exec -it pg bash

# restore the database
- cd mydata
- pg_restore -U postgres -d dvdrental dvdrental.tar

# review the data in postgres
- psql -h localhost -U postgres
- \c dvdrental;
- select * from actor;
```

- Run pgadmin4 into docker container:
```
# run the container
docker run --rm -it -p 5151:80 --name=pnl -e "PGADMIN_DEFAULT_EMAIL=m@mod.com" -e "PGADMIN_DEFAULT_PASSWORD=admin" dpage/pgadmin4

# open pgadmin4 in browser
http://localhost:5151

# login
email: m@mod.com
pass: admin

# register - server
- Name: {pick any name}
- Host: 172.17.0.1 {or use docker inspect}
  $ docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' <container-ID>
- port: 5454
- DB=dvdrental
- user: postgres
- pass: postgres
```

- Run the application.
```
$ vi config.json
  {
    "AppName"               :   "Application",
    "Module"                :   "github.com/Application",
    "DBConnectionString"    :   "host=127.0.0.1 port=5432 user=postgres password=postgres dbname=dvdrental sslmode=disable",
    "DBEngine"              :   "postgres",
    "WebAddress"            :   "0.0.0.0",
    "WebPort"               :   7070
  }
  
$ modelgen 
$ cd Application/
  ├──> Application
    ├──> config.json
    ├──> go.mod 
    ├──> main.go
    ├──> Dockerfile
    db
    │	├──> database.go
    api
    │	├──> ActorAPI.go
    │ ...
    bll
    │	├──> ActorBLL.go
    │ ...
    dal
    │	├──> ActorDAL.go
    │ ...
    dto
    │	├──> ActorDTO.go
    │ ...
    test
    │	├──> Actor_test.go
    │ ...
    │	├──> test.json
    ─────────────────────────────

```
## Docker 
```
# build docker image
docker build -t concept-api .
```
```
# run docker container 
docker run --rm -it -p 7070:7070 --env-file .env --network host concept-api
```

## Concept
- in [![standard](https://github.com/tarekbadrshalaan/modelgen/tree/master/modelgen-concept/web-api)](https://github.com/tarekbadrshalaan/modelgen/tree/master/modelgen-concept/web-api) you can find the expected result of the generator. 

## Contributing

PRs accepted.


## License
[![License: MIT](https://img.shields.io/badge/License-MIT-ff69b4.svg)](https://opensource.org/licenses/MIT)
