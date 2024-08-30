package routes

import (
	"net/http"
	"api/src/middlewares"
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
	routes = append(routes, loginRoute)
	routes = append(routes, googleapi...)

	for _, route :=  range routes{
		if route.RequestAuthentication {
			r.HandleFunc(route.URI, 
				middlewares.Logger(middlewares.Authenticate(route.Function)),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middlewares.Logger(route.Function)).Methods(route.Method)
		}
		
	}

	return r
}