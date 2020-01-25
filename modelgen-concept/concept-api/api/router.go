package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
)

// ConfigRouter : configure endpoints in the server.
func ConfigRouter() http.Handler {
	router := httprouter.New()

	configActorsRouter(router)
	configAddressesRouter(router)
	configCategoriesRouter(router)
	configCitiesRouter(router)
	configCountriesRouter(router)
	configCustomersRouter(router)
	configFilmsRouter(router)
	configInventoriesRouter(router)
	configLanguagesRouter(router)
	configPaymentsRouter(router)
	configRentalsRouter(router)
	configStaffsRouter(router)
	configStoresRouter(router)
	

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
