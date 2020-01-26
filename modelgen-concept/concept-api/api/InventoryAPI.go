package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configInventoriesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/inventories", handle: getAllInventories})
	*routes = append(*routes, route{method: "POST", path:"/inventories", handle: postInventories})
	*routes = append(*routes, route{method: "PUT", path:"/inventories", handle: putInventories})
	*routes = append(*routes, route{method: "GET", path:"/inventories/:id", handle: getInventories})
	*routes = append(*routes, route{method: "DELETE", path:"/inventories/:id", handle: deleteInventories})
}

func getAllInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventories, err := bll.GetAllInventories()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, inventories, http.StatusOK)
}

func getInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertInventoryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	inventory, err := bll.GetInventory(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find inventory (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, inventory, http.StatusOK)
}

func postInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventory := &dto.InventoryDTO{}
	if err := readJSON(r, inventory); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateInventory(inventory)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventory := &dto.InventoryDTO{}
	if err := readJSON(r, inventory); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateInventory(inventory)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertInventoryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteInventory(id)
	if err != nil {
		msg := fmt.Errorf("Inventory with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
