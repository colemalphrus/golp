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
	fs := http.FileServer(http.Dir("public/"))
	r.PathPrefix("/public/").Handler(http.StripPrefix("/public/", fs))

	http.ListenAndServe(port, r)
}