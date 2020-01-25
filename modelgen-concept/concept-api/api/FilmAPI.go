package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, films)
}

func getFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertFilmID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	film, err := bll.GetFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, film)
}

func postFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	film := &dto.FilmDTO{}
	if err := readJSON(r, film); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateFilm(film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	film := &dto.FilmDTO{}
	if err := readJSON(r, film); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateFilm(film)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func deleteFilms(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertFilmID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteFilm(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
