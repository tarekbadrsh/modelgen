package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configCountriesRouter(router *httprouter.Router) {
	router.GET("/countries", getAllCountries)
	router.POST("/countries", postCountries)
	router.PUT("/countries", putCountries)
	router.GET("/countries/:id", getCountries)
	router.DELETE("/countries/:id", deleteCountries)
}

func getAllCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	countries, err := bll.GetAllCountries()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, countries)
}

func getCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCountryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	country, err := bll.GetCountry(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, country)
}


func postCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	country := &dto.CountryDTO{}
	if err := readJSON(r, country); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCountry(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	country := &dto.CountryDTO{}
	if err := readJSON(r, country); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCountry(country)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteCountries(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertCountryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteCountry(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

