package main

import (
  "net/http"
  "log"
  "github.com/julienschmidt/httprouter"
  c "./controllers"
)

func main()  {
  router := httprouter.New()
  router.GET("/", c.Index)
  router.NotFound = http.StripPrefix("/public/", http.FileServer(http.Dir("./static/")))
  log.Fatal(http.ListenAndServe(":8000", router))
}
