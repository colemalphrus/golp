package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

func About(w http.ResponseWriter, r *http.Request){
	data := map[string]string{}

	templateFile, _ := filepath.Abs("views/about.html.mustache")
	view, _ := mustache.RenderFile(templateFile, data)

	fmt.Fprintf(w, view)
}