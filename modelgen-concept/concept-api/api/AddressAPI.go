package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, addresses)
}

func getAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertAddressID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	address, err := bll.GetAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, address)
}

func postAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dto.AddressDTO{}
	if err := readJSON(r, address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	address := &dto.AddressDTO{}
	if err := readJSON(r, address); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateAddress(address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func deleteAddresses(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertAddressID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteAddress(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
