package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configCategoriesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/categories", handle: getAllCategories})
	*routes = append(*routes, route{method: "POST", path:"/categories", handle: postCategories})
	*routes = append(*routes, route{method: "PUT", path:"/categories", handle: putCategories})
	*routes = append(*routes, route{method: "GET", path:"/categories/:id", handle: getCategories})
	*routes = append(*routes, route{method: "DELETE", path:"/categories/:id", handle: deleteCategories})
}

func getAllCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categories, err := bll.GetAllCategories()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, categories, http.StatusOK)
}

func getCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCategoryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	category, err := bll.GetCategory(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find category (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, category, http.StatusOK)
}

func postCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCategory(category)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCategory(category)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCategoryID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteCategory(id)
	if err != nil {
		msg := fmt.Errorf("Category with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
