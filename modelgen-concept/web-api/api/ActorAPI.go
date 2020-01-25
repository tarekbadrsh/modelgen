package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/web-api/bll"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)

func configActorsRouter(router *httprouter.Router) {
	router.GET("/actors", getAllActors)
	router.POST("/actors", postActors)
	router.PUT("/actors", putActors)
	router.GET("/actors/:id", getActors)
	router.DELETE("/actors/:id", deleteActors)
}

func getAllActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actors, err := bll.GetAllActors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, actors)
}

func getActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertActorID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	actor, err := bll.GetActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, actor)
}


func postActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actor := &dto.ActorDTO{}
	if err := readJSON(r, actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateActor(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actor := &dto.ActorDTO{}
	if err := readJSON(r, actor); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateActor(actor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertActorID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteActor(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

