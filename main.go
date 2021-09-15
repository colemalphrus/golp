package main

import (
	"net/http"

	"github.com/colemalphrus/golp/controllers"
)

func main (){
	http.HandleFunc("/about", controllers.About)
	http.HandleFunc("/", controllers.Homepage)
	http.ListenAndServe(":8080", nil)
}