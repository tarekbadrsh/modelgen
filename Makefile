.PHONY: help build run

# set default goal to help
.DEFAULT_GOAL := help


### * make help                                             		Print this help
help: Makefile
	@sed -n 's/^###//p' $<

### * make build								Build modelgen binary and run modelgen to generate concept app 
build:
	- go build .
	- ./modelgen

### * make run								run concept-api (concept) binary 
run: build
	- cd modelgen-concept/concept-api && go build .
	- cd modelgen-concept/concept-api && ./concept-api
