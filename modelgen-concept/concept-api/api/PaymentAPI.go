package api

import (
	"fmt"
	"net/http"

	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/bll"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/dto"
	"github.com/tarekbadrshalaan/modelgen/modelgen-concept/concept-api/logger"
	
	"github.com/julienschmidt/httprouter"
)

func configPaymentsRouter(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/payments", handle: getAllPayments})
	*routes = append(*routes, route{method: "POST", path:"/payments", handle: postPayments})
	*routes = append(*routes, route{method: "PUT", path:"/payments", handle: putPayments})
	*routes = append(*routes, route{method: "GET", path:"/payments/:id", handle: getPayments})
	*routes = append(*routes, route{method: "DELETE", path:"/payments/:id", handle: deletePayments})
}

func getAllPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payments, err := bll.GetAllPayments()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, payments, http.StatusOK)
}

func getPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertPaymentID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	

	payment, err := bll.GetPayment(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find payment (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, payment, http.StatusOK)
}

func postPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payment := &dto.PaymentDTO{}
	if err := readJSON(r, payment); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.CreatePayment(payment)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func putPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payment := &dto.PaymentDTO{}
	if err := readJSON(r, payment); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.UpdatePayment(payment)
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func deletePayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	requestID := ps.ByName("id")
	id, err := bll.ConvertPaymentID(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be int32; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	
	
	err = bll.DeletePayment(id)
	if err != nil {
		msg := fmt.Errorf("Payment with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
