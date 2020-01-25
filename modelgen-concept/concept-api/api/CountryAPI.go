package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configCountriesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/countries", handle: getAllCountries})
	*routes = append(*routes, route{method: "POST", path:"/countries", handle: postCountries})
	*routes = append(*routes, route{method: "PUT", path:"/countries", handle: putCountries})
	*routes = append(*routes, route{method: "GET", path:"/countries/:id", handle: getCountries})
	*routes = append(*routes, route{method: "DELETE", path:"/countries/:id", handle: deleteCountries})
}

func getAllCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	countries, err := bll.GetAllCountries()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, countries, http.StatusOK)
}

func getCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCountryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	country, err := bll.GetCountry(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find country (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, country, http.StatusOK)
}

func postCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	country := &dto.CountryDTO{}
	if err := readJSON(r, country); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCountry(country)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	country := &dto.CountryDTO{}
	if err := readJSON(r, country); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCountry(country)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCountryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteCountry(id)
	if err != nil {
		msg := fmt.Errorf("Country with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
