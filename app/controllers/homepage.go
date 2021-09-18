package controllers

import (
	"fmt"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/colemalphrus/golp/golp"
)

func Homepage(w http.ResponseWriter, r *http.Request){

	data := map[string]string{
		"title": "Golp",
		"header": "Welcome to GoLP",
	}

	templateFile := golp.Template("app/views/", "homepage.html.mustache")
	view, _ := mustache.RenderFile(templateFile, data)

	fmt.Fprintf(w, view)
} 