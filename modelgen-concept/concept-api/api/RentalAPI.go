package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configRentalsRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/rentals", handle: getAllRentals})
	*routes = append(*routes, route{method: "POST", path:"/rentals", handle: postRentals})
	*routes = append(*routes, route{method: "PUT", path:"/rentals", handle: putRentals})
	*routes = append(*routes, route{method: "GET", path:"/rentals/:id", handle: getRentals})
	*routes = append(*routes, route{method: "DELETE", path:"/rentals/:id", handle: deleteRentals})
}

func getAllRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rentals, err := bll.GetAllRentals()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, rentals, http.StatusOK)
}

func getRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertRentalID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	rental, err := bll.GetRental(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find rental (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, rental, http.StatusOK)
}

func postRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rental := &dto.RentalDTO{}
	if err := readJSON(r, rental); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateRental(rental)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	rental := &dto.RentalDTO{}
	if err := readJSON(r, rental); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateRental(rental)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteRentals(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertRentalID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteRental(id)
	if err != nil {
		msg := fmt.Errorf("Rental with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
