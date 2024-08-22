package routes

import (
	"api/src/controllers"
	"net/http"
)

var googleapi = []Routes{
	{
		URI:                   "/googlebooks",
		Method:                http.MethodGet,
		Function:              controllers.SearchGoogleBooks,
		RequestAuthentication: false,
	},
	{
		URI:                   "/googlebooks/{title}",
		Method:                http.MethodGet,
		Function:              controllers.SearchGoogleBooksByTitle,
		RequestAuthentication: false,
	},
}
