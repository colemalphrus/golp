package controllers

import (
	"fmt"
	"net/http"

	"github.com/cbroglie/mustache"
	"github.com/colemalphrus/golp/app_svelte/config"
	"github.com/colemalphrus/golp/golp"
)

func App(w http.ResponseWriter, r *http.Request){
	data := map[string]string{}

	templateFile := golp.Template(config.TemplateDir, "app.html.mustache")
	view, _ := mustache.RenderFile(templateFile, data)

	fmt.Fprintf(w, view)
}