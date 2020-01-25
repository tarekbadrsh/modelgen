package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configAddressesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/addresses", handle: getAllAddresses})
	*routes = append(*routes, route{method: "POST", path:"/addresses", handle: postAddresses})
	*routes = append(*routes, route{method: "PUT", path:"/addresses", handle: putAddresses})
	*routes = append(*routes, route{method: "GET", path:"/addresses/:id", handle: getAddresses})
	*routes = append(*routes, route{method: "DELETE", path:"/addresses/:id", handle: deleteAddresses})
}

func getAllAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	addresses, err := bll.GetAllAddresses()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, addresses, http.StatusOK)
}

func getAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertAddressID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	address, err := bll.GetAddress(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find address (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, address, http.StatusOK)
}

func postAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dto.AddressDTO{}
	if err := readJSON(r, address); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateAddress(address)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dto.AddressDTO{}
	if err := readJSON(r, address); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateAddress(address)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertAddressID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteAddress(id)
	if err != nil {
		msg := fmt.Errorf("Address with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
