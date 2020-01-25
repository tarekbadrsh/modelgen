.PHONY: help build run
.PHONY: help build-api run-api
.PHONY: build-docker run-docker

# set default goal to help
.DEFAULT_GOAL := help


### * make help                                             		Print this help
help: Makefile
	@sed -n 's/^###//p' $<

### * make build								Build modelgen binary
build:
	- go build .

### * make run								run modelgen to generate concept app 
run:build
	- ./modelgen

### * make build-api								build concept-api (concept) binary 
build-api:run
	- cd modelgen-concept/concept-api && go build .

### * make run-api								run concept-api (concept) binary 
run-api: build-api
	- cd modelgen-concept/concept-api && ./concept-api

### * make build-docker							Build modelgen docker-image (concept-api)
build-docker:build-api
	- cd modelgen-concept/concept-api && docker build -t concept-api .

### * make run-docker							run modelgen docker-image (concept-api)
run-docker:build-docker
	- cd modelgen-concept/concept-api && docker run --rm -it -p 7070:7070 --env-file .env --network host concept-api 
