package templates

// apiTmpl : template of API
var apiTmpl = `package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"{{.Module}}/bll"
	"{{.Module}}/dto"
)

func config{{pluralize .StructName}}Router(router *httprouter.Router) {
	router.GET("/{{pluralizeLower .StructName}}", getAll{{pluralize .StructName}})
	router.POST("/{{pluralizeLower .StructName}}", post{{pluralize .StructName}})
	router.PUT("/{{pluralizeLower .StructName}}", put{{pluralize .StructName}})
	router.GET("/{{pluralizeLower .StructName}}/:id", get{{pluralize .StructName}})
	router.DELETE("/{{pluralizeLower .StructName}}/:id", delete{{pluralize .StructName}})
}

func getAll{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{pluralizeLower .StructName}}, err := bll.GetAll{{pluralize .StructName}}()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	writeJSON(w, {{pluralizeLower .StructName}})
}

func get{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{ if (ne .IDType "string")}}
	id, err := bll.Convert{{.IDName}}(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be {{.IDType}}", http.StatusBadRequest)
		return
	}
	{{ else }}
	id := ps.ByName("id")
	{{ end }}

	{{toLower .StructName}}, err := bll.Get{{.StructName}}(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	writeJSON(w, {{toLower .StructName}})
}


func post{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{toLower .StructName}} := &dto.{{DTO .StructName}}{}
	if err := readJSON(r, {{toLower .StructName}}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.Create{{.StructName}}({{toLower .StructName}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}

func put{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{toLower .StructName}} := &dto.{{DTO .StructName}}{}
	if err := readJSON(r, {{toLower .StructName}}); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result, err := bll.Update{{.StructName}}({{toLower .StructName}})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	writeJSON(w, result)
}


func delete{{pluralize .StructName}}(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	{{ if (eq .IDType "string")}}	
	id := ps.ByName("id")
	{{ else }}
	id, err := bll.Convert{{.IDName}}(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Error: parameter (id) should be {{.IDType}}", http.StatusBadRequest)
		return
	}
	{{ end }}

	err = bll.Delete{{.StructName}}(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

`
