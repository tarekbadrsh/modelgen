package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/tarekbadrshalaan/modelgen/web-api/bll"
	"github.com/tarekbadrshalaan/modelgen/web-api/dto"
)

func configPaymentsRouter(router *httprouter.Router) {
	router.GET("/payments", getAllPayments)
	router.POST("/payments", postPayments)
	router.PUT("/payments", putPayments)
	router.GET("/payments/:id", getPayments)
	router.DELETE("/payments/:id", deletePayments)
}

func getAllPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payments, err := bll.GetAllPayments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, payments)
}

func getPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertPaymentID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	payment, err := bll.GetPayment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, payment)
}


func postPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payment := &dto.PaymentDTO{}
	if err := readJSON(r, payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.CreatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func putPayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	payment := &dto.PaymentDTO{}
	if err := readJSON(r, payment); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.UpdatePayment(payment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func deletePayments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	id, err := bll.ConvertPaymentID(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be int32", http.StatusBadRequest)
		return
	}
	

	err = bll.DeletePayment(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

