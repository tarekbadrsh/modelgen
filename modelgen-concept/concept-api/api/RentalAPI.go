package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configRentalsRouter(router *httprouter.Router) {
	router.GET("/rentals", getAllRentals)
	router.POST("/rentals", postRentals)
	router.PUT("/rentals", putRentals)
	router.GET("/rentals/:id", getRentals)
	router.DELETE("/rentals/:id", deleteRentals)
}

func getAllRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rentals, err := bll.GetAllRentals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, rentals)
}

func getRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertRentalID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	rental, err := bll.GetRental(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, rental)
}


func postRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rental := &dto.RentalDTO{}
	if err := readJSON(r, rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateRental(rental)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rental := &dto.RentalDTO{}
	if err := readJSON(r, rental); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateRental(rental)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertRentalID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteRental(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

