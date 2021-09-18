package golp

import "net/http"

type ControllerFunction func(http.ResponseWriter, *http.Request)

type Route struct {
	slug       string
	controller ControllerFunction
}

func ParseRoutes() {

}