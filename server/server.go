package server

import (
	"net/http"

	"github.com/gorilla/mux"
)


func Serve(port string){

	r := mux.NewRouter()

	for _, v := range Routes {
		r.HandleFunc(
			v.slug,
			v.controller,
		)
	}

	http.ListenAndServe(port, r)
}