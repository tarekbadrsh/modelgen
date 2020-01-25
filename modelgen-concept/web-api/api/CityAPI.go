package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/web-api/bll"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)

func configCitiesRouter(router *httprouter.Router) {
	router.GET("/cities", getAllCities)
	router.POST("/cities", postCities)
	router.PUT("/cities", putCities)
	router.GET("/cities/:id", getCities)
	router.DELETE("/cities/:id", deleteCities)
}

func getAllCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cities, err := bll.GetAllCities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, cities)
}

func getCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCityID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	city, err := bll.GetCity(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, city)
}


func postCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := &dto.CityDTO{}
	if err := readJSON(r, city); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	city := &dto.CityDTO{}
	if err := readJSON(r, city); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCity(city)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteCities(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCityID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteCity(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

