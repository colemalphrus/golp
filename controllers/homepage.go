package controllers

import (
	"fmt"
	"net/http"
)

func Homepage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Welcome to my website!")
}