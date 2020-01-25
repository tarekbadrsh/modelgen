package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/web-api/dto"
)

func configStaffsRouter(router *httprouter.Router) {
	router.GET("/staffs", getAllStaffs)
	router.POST("/staffs", postStaffs)
	router.PUT("/staffs", putStaffs)
	router.GET("/staffs/:id", getStaffs)
	router.DELETE("/staffs/:id", deleteStaffs)
}

func getAllStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staffs, err := bll.GetAllStaffs()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, staffs)
}

func getStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertStaffID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	staff, err := bll.GetStaff(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, staff)
}


func postStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staff := &dto.StaffDTO{}
	if err := readJSON(r, staff); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreateStaff(staff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staff := &dto.StaffDTO{}
	if err := readJSON(r, staff); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateStaff(staff)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deleteStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertStaffID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeleteStaff(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

