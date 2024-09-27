package routes

import (
	"api/src/controllers"
	"net/http"
)

var files = []Routes{
	{
		URI:                   "/static/{file}",
		Method:                http.MethodGet,
		Function:              controllers.ServeStaticFiles,
		RequestAuthentication: false,
	},
}