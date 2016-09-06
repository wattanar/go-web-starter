package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
  c "github.com/wattanar/go-web-starter/controllers"
)

func main()  {

  r := httprouter.New()

  r.GET("/", c.Index)
  
  r.NotFound = http.StripPrefix("/static/", http.FileServer(http.Dir("./assets/")))
  
  log.Fatal(http.ListenAndServe(":8080", r))
}