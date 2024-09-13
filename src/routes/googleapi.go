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
	{
		URI:                   "/googlebooks/add",
		Method:                http.MethodPost,
		Function:              controllers.AddGoogleBook,
		RequestAuthentication: false,
	},
	{
		URI:                   "/clean_database",
		Method:                http.MethodGet,
		Function:              controllers.CleanDatabase,
		RequestAuthentication: false,
	},
}
