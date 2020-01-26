package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
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
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, stores, http.StatusOK)
}

func getStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertStoreID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	store, err := bll.GetStore(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find store (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, store, http.StatusOK)
}

func postStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	store := &dto.StoreDTO{}
	if err := readJSON(r, store); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateStore(store)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	store := &dto.StoreDTO{}
	if err := readJSON(r, store); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateStore(store)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteStores(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertStoreID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteStore(id)
	if err != nil {
		msg := fmt.Errorf("Store with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
