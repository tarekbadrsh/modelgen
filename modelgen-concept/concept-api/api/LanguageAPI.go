package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configLanguagesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/languages", handle: getAllLanguages})
	*routes = append(*routes, route{method: "POST", path:"/languages", handle: postLanguages})
	*routes = append(*routes, route{method: "PUT", path:"/languages", handle: putLanguages})
	*routes = append(*routes, route{method: "GET", path:"/languages/:id", handle: getLanguages})
	*routes = append(*routes, route{method: "DELETE", path:"/languages/:id", handle: deleteLanguages})
}

func getAllLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	languages, err := bll.GetAllLanguages()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, languages, http.StatusOK)
}

func getLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertLanguageID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	language, err := bll.GetLanguage(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find language (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, language, http.StatusOK)
}

func postLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	language := &dto.LanguageDTO{}
	if err := readJSON(r, language); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateLanguage(language)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	language := &dto.LanguageDTO{}
	if err := readJSON(r, language); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateLanguage(language)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertLanguageID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteLanguage(id)
	if err != nil {
		msg := fmt.Errorf("Language with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
