package main

import (
	"html/template"
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	t, _ := template.ParseFiles("views/base.tmpl", "views/index.tmpl")
	t.Execute(w, nil)
}

func main()  {
	router := httprouter.New()
	router.GET("/", Index)
	router.NotFound = http.StripPrefix("/public/", http.FileServer(http.Dir("./static/")))
	log.Fatal(http.ListenAndServe(":8000", router))
}
