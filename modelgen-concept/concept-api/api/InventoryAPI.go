package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configInventoriesRouter(router *httprouter.Router) {
	router.GET("/inventories", getAllInventories)
	router.POST("/inventories", postInventories)
	router.PUT("/inventories", putInventories)
	router.GET("/inventories/:id", getInventories)
	router.DELETE("/inventories/:id", deleteInventories)
}

func getAllInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventories, err := bll.GetAllInventories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, inventories)
}

func getInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertInventoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	inventory, err := bll.GetInventory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, inventory)
}


func postInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventory := &dto.InventoryDTO{}
	if err := readJSON(r, inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateInventory(inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	inventory := &dto.InventoryDTO{}
	if err := readJSON(r, inventory); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateInventory(inventory)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteInventories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertInventoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteInventory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

