package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configActorsRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/actors", handle: getAllActors})
	*routes = append(*routes, route{method: "POST", path:"/actors", handle: postActors})
	*routes = append(*routes, route{method: "PUT", path:"/actors", handle: putActors})
	*routes = append(*routes, route{method: "GET", path:"/actors/:id", handle: getActors})
	*routes = append(*routes, route{method: "DELETE", path:"/actors/:id", handle: deleteActors})
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
