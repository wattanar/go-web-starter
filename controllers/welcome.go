package controllers

import (
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("views/base.tmpl", "views/index.tmpl")
	t.Execute(w, nil)
}