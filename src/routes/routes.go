package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI      				string
	Method   				string
	Function 				func(http.ResponseWriter, *http.Request)
	RequestAuthentication 	bool
}

func Config(r *mux.Router) *mux.Router{

	routes := usersRoutes
	routes = append(routes, googleapi...)

	for _, route :=  range routes{
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}