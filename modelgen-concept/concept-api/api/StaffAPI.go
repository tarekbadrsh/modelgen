package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configStaffsRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/staffs", handle: getAllStaffs})
	*routes = append(*routes, route{method: "POST", path:"/staffs", handle: postStaffs})
	*routes = append(*routes, route{method: "PUT", path:"/staffs", handle: putStaffs})
	*routes = append(*routes, route{method: "GET", path:"/staffs/:id", handle: getStaffs})
	*routes = append(*routes, route{method: "DELETE", path:"/staffs/:id", handle: deleteStaffs})
}

func getAllStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staffs, err := bll.GetAllStaffs()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, staffs, http.StatusOK)
}

func getStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertStaffID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	staff, err := bll.GetStaff(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find staff (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, staff, http.StatusOK)
}

func postStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staff := &dto.StaffDTO{}
	if err := readJSON(r, staff); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateStaff(staff)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	staff := &dto.StaffDTO{}
	if err := readJSON(r, staff); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateStaff(staff)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteStaffs(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertStaffID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteStaff(id)
	if err != nil {
		msg := fmt.Errorf("Staff with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
