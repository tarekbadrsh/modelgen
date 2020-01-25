package templates

// apiRouterTmpl : template of APIRouterTmpl
var apiRouterTmpl = `package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
)

type route struct {
	method string            //HTTP method
	path   string            //url endpoint
	handle httprouter.Handle //Controller function which dispatches the right HTML page and/or data for each route
}

// configRouter : configure endpoints in the server.
func configRouter() *[]route {
	routes := &[]route{}
	{{ range $StructName, $value := .}}config{{pluralize $StructName}}Router(routes)
	{{ end }}
	return routes
}

// NewRouter :creates a new router instance and iterate through all the Routes to get each’s
// Route’s Method, Pattern and Handle and registers a new request handle.
func NewRouter() http.Handler {
	routes := configRouter()
	router := httprouter.New()
	for _, route := range *routes {
		router.Handle(route.method, route.path, logmid(route.handle))
	}
	return router
}

func writeJSON(w http.ResponseWriter, v interface{}) {
	data, _ := json.Marshal(v)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	// allow cross domain AJAX requests
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

	w.Write(data)
}

func readJSON(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(buf, v)
}

// logmid : logging midleware
func logmid(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger.Infof("[%s] on: %s", r.Method, r.URL)
		next(w, r, ps)
	}
}
`
