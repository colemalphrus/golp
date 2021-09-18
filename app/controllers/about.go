package controllers

import (
	"fmt"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/colemalphrus/golp/golp"
)

func About(w http.ResponseWriter, r *http.Request){
	data := map[string]string{}

	templateFile := golp.Template("app/views/", "about.html.mustache")
	view, _ := mustache.RenderFile(templateFile, data)

	fmt.Fprintf(w, view)
}