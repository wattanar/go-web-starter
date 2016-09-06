package controllers

import (
	"net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

func Index(res http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("views/base.tmpl", "views/index.tmpl")
	t.Execute(res, nil)
}