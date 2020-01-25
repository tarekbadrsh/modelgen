package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
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
