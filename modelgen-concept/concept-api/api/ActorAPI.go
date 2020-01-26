package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
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
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, actors, http.StatusOK)
}

func getActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertActorID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	actor, err := bll.GetActor(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find actor (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, actor, http.StatusOK)
}

func postActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actor := &dto.ActorDTO{}
	if err := readJSON(r, actor); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateActor(actor)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	actor := &dto.ActorDTO{}
	if err := readJSON(r, actor); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateActor(actor)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteActors(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertActorID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteActor(id)
	if err != nil {
		msg := fmt.Errorf("Actor with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
