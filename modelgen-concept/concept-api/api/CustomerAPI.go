package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configCustomersRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/customers", handle: getAllCustomers})
	*routes = append(*routes, route{method: "POST", path:"/customers", handle: postCustomers})
	*routes = append(*routes, route{method: "PUT", path:"/customers", handle: putCustomers})
	*routes = append(*routes, route{method: "GET", path:"/customers/:id", handle: getCustomers})
	*routes = append(*routes, route{method: "DELETE", path:"/customers/:id", handle: deleteCustomers})
}

func getAllCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customers, err := bll.GetAllCustomers()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, customers, http.StatusOK)
}

func getCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCustomerID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	customer, err := bll.GetCustomer(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find customer (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, customer, http.StatusOK)
}

func postCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customer := &dto.CustomerDTO{}
	if err := readJSON(r, customer); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreateCustomer(customer)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	customer := &dto.CustomerDTO{}
	if err := readJSON(r, customer); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdateCustomer(customer)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deleteCustomers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertCustomerID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeleteCustomer(id)
	if err != nil {
		msg := fmt.Errorf("Customer with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
