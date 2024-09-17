package routes

import (
	"api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Routes struct {
	URI                   string
	Method                string
	Function              func(http.ResponseWriter, *http.Request)
	RequestAuthentication bool
}

func Config(r *mux.Router) *mux.Router {

	routes := usersRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, refresh_token...)
	routes = append(routes, files...)
	routes = append(routes, reviews...)
	routes = append(routes, books...)
	routes = append(routes, googleapi...)

	for _, route := range routes {
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
