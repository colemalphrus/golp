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
	fs := http.FileServer(http.Dir("static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	http.ListenAndServe(port, r)
}