package server

import (
	"net/http"

	"github.com/colemalphrus/golp/app/controllers"
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
}