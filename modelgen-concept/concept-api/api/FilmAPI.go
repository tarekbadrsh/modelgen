package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configFilmsRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/films", handle: getAllFilms})
	*routes = append(*routes, route{method: "POST", path:"/films", handle: postFilms})
	*routes = append(*routes, route{method: "PUT", path:"/films", handle: putFilms})
	*routes = append(*routes, route{method: "GET", path:"/films/:id", handle: getFilms})
	*routes = append(*routes, route{method: "DELETE", path:"/films/:id", handle: deleteFilms})
}

func getAllFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	films, err := bll.GetAllFilms()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, films, http.StatusOK)
}

func getFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertFilmID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	film, err := bll.GetFilm(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find film (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, film, http.StatusOK)
}

func postFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	film := &dto.FilmDTO{}
	if err := readJSON(r, film); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateFilm(film)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	film := &dto.FilmDTO{}
	if err := readJSON(r, film); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateFilm(film)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertFilmID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteFilm(id)
	if err != nil {
		msg := fmt.Errorf("Film with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
