package controllers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func User(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	
	fmt.Fprintf(w, vars["id"])
}