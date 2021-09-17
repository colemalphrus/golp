package controllers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/cbroglie/mustache"
)

func Homepage(w http.ResponseWriter, r *http.Request){

	templateFile, _ := filepath.Abs("views/homepage.html.mustache")
	view, _ := mustache.RenderFile(templateFile, map[string]string{"name": "cole malphrus"})

	fmt.Fprintf(w, view)
} 