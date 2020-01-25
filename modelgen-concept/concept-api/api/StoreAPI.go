package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configStoresRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/stores", handle: getAllStores})
	*routes = append(*routes, route{method: "POST", path:"/stores", handle: postStores})
	*routes = append(*routes, route{method: "PUT", path:"/stores", handle: putStores})
	*routes = append(*routes, route{method: "GET", path:"/stores/:id", handle: getStores})
	*routes = append(*routes, route{method: "DELETE", path:"/stores/:id", handle: deleteStores})
}

func getAllStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	stores, err := bll.GetAllStores()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, stores)
}

func getStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertStoreID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	store, err := bll.GetStore(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, store)
}

func postStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	store := &dto.StoreDTO{}
	if err := readJSON(r, store); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateStore(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	store := &dto.StoreDTO{}
	if err := readJSON(r, store); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateStore(store)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func deleteStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertStoreID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteStore(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
