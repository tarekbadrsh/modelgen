package templates

// apiTmpl : template of API
var apiTmpl = `package api

import (
	"fmt"
	"net/http"

	"{{.Module}}/bll"
	"{{.Module}}/dto"
	"{{.Module}}/logger"
	
	"github.com/julienschmidt/httprouter"
)

func config{{pluralize .StructName}}Router(routes *[]route) {
	*routes = append(*routes, route{method: "GET", path:"/{{pluralizeLower .StructName}}", handle: getAll{{pluralize .StructName}}})
	*routes = append(*routes, route{method: "POST", path:"/{{pluralizeLower .StructName}}", handle: post{{pluralize .StructName}}})
	*routes = append(*routes, route{method: "PUT", path:"/{{pluralizeLower .StructName}}", handle: put{{pluralize .StructName}}})
	*routes = append(*routes, route{method: "GET", path:"/{{pluralizeLower .StructName}}/:id", handle: get{{pluralize .StructName}}})
	*routes = append(*routes, route{method: "DELETE", path:"/{{pluralizeLower .StructName}}/:id", handle: delete{{pluralize .StructName}}})
}

func getAll{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{pluralizeLower .StructName}}, err := bll.GetAll{{pluralize .StructName}}()
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, {{pluralizeLower .StructName}}, http.StatusOK)
}

func get{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{ if (eq .IDType "string")}}	
	id := ps.ByName("id")
	{{ else }}
	requestID := ps.ByName("id")
	id, err := bll.Convert{{.IDName}}(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be {{.IDType}}; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	{{ end }}

	{{toLower .StructName}}, err := bll.Get{{.StructName}}(id)
	if err != nil {
		msg := fmt.Errorf("Can’t find {{toLower .StructName}} (%v); err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return
	}
	writeResponseJSON(w, {{toLower .StructName}}, http.StatusOK)
}

func post{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{toLower .StructName}} := &dto.{{DTO .StructName}}{}
	if err := readJSON(r, {{toLower .StructName}}); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.Create{{.StructName}}({{toLower .StructName}})
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusInternalServerError)
		return
	}
	writeResponseJSON(w, result, http.StatusCreated)
}

func put{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{toLower .StructName}} := &dto.{{DTO .StructName}}{}
	if err := readJSON(r, {{toLower .StructName}}); err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}

	result, err := bll.Update{{.StructName}}({{toLower .StructName}})
	if err != nil {
		logger.Error(err)
		writeResponseError(w, err, http.StatusBadRequest)
		return
	}
	writeResponseJSON(w, result, http.StatusOK)
}

func delete{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{ if (eq .IDType "string")}}	
	id := ps.ByName("id")
	{{ else }}
	requestID := ps.ByName("id")
	id, err := bll.Convert{{.IDName}}(requestID)
	if err != nil {
		msg := fmt.Errorf("Error: parameter (id) should be {{.IDType}}; Id=%v; err (%v)", requestID, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusBadRequest)
		return
	}
	{{ end }}
	
	err = bll.Delete{{.StructName}}(id)
	if err != nil {
		msg := fmt.Errorf("{{.StructName}} with id (%v) does not exist; err (%v)", id, err)
		logger.Error(msg)
		writeResponseError(w, msg, http.StatusNotFound)
		return

	}
	writeResponseJSON(w, true, http.StatusOK)
}
`
