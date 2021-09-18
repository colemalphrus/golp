package server

import (
	"net/http"

	"github.com/colemalphrus/golp/app/controllers"
	svelte_c "github.com/colemalphrus/golp/app_svelte/controllers"
)

type controllerFunction func(http.ResponseWriter, *http.Request)

type route struct {
	slug       string
	controller controllerFunction
}

var Routes = []route{
	{slug: "/", controller: controllers.Homepage},
	{slug: "/about", controller: controllers.About},
	{slug: "/user/{id}", controller: controllers.User},
	{slug: "/app", controller: svelte_c.App},
}