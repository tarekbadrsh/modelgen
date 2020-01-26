package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configCitiesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/cities", handle: getAllCities})
	*routes = append(*routes, route{method: "POST", path:"/cities", handle: postCities})
	*routes = append(*routes, route{method: "PUT", path:"/cities", handle: putCities})
	*routes = append(*routes, route{method: "GET", path:"/cities/:id", handle: getCities})
	*routes = append(*routes, route{method: "DELETE", path:"/cities/:id", handle: deleteCities})
}

func getAllCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cities, err := bll.GetAllCities()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, cities, http.StatusOK)
}

func getCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCityID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	city, err := bll.GetCity(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find city (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, city, http.StatusOK)
}

func postCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := &dto.CityDTO{}
	if err := readJSON(r, city); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCity(city)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := &dto.CityDTO{}
	if err := readJSON(r, city); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCity(city)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCityID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteCity(id)
	if err != nil {
		msg := fmt.Errorf("City with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
