package main

import (
  "fmt"
  "net/http"
  "strconv"

  "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Haga una petición GET con numeros a esta ruta.")
}


func Multiple(w http.ResponseWriter, r *http.Request)  {
  vars := mux.Vars(r)
  Id := vars["Id"]
  R, err := strconv.Atoi(Id)
  if err != nil {

  }
  R= R * 2
  fmt.Fprintln(w, "Su numero es:", R)
}
