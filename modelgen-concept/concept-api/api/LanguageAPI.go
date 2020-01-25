package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configLanguagesRouter(router *httprouter.Router) {
	router.GET("/languages", getAllLanguages)
	router.POST("/languages", postLanguages)
	router.PUT("/languages", putLanguages)
	router.GET("/languages/:id", getLanguages)
	router.DELETE("/languages/:id", deleteLanguages)
}

func getAllLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	languages, err := bll.GetAllLanguages()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, languages)
}

func getLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertLanguageID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	language, err := bll.GetLanguage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, language)
}


func postLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	language := &dto.LanguageDTO{}
	if err := readJSON(r, language); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateLanguage(language)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	language := &dto.LanguageDTO{}
	if err := readJSON(r, language); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateLanguage(language)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteLanguages(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertLanguageID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteLanguage(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

