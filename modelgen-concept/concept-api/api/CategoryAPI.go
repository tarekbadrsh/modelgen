package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
)

func configCategoriesRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path: "/categories", handle: getAllCategories})
	*routes = append(*routes, route{method: "POST", path: "/categories", handle: postCategories})
	*routes = append(*routes, route{method: "PUT", path: "/categories", handle: putCategories})
	*routes = append(*routes, route{method: "GET", path: "/categories/:id", handle: getCategories})
	*routes = append(*routes, route{method: "DELETE", path: "/categories/:id", handle: deleteCategories})
}

func getAllCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	categories, err := bll.GetAllCategories()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, categories)
}

func getCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := bll.ConvertCategoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}

	category, err := bll.GetCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, category)
}

func postCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	category := &dto.CategoryDTO{}
	if err := readJSON(r, category); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCategory(category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func deleteCategories(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := bll.ConvertCategoryID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}

	err = bll.DeleteCategory(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
