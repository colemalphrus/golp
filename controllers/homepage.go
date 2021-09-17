package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

func Homepage(w http.ResponseWriter, r *http.Request){

	data := map[string]string{
		"title": "Golp",
		"header": "Welcome to GoLP",
	}

	templateFile, _ := filepath.Abs("views/homepage.html.mustache")
	view, _ := mustache.RenderFile(templateFile, data)

	fmt.Fprintf(w, view)
} 