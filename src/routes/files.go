package routes

import (
	"api/src/controllers"
	"net/http"
)

var files = []Routes{
	{
		URI:                   "/files/{book_id}",
		Method:                http.MethodPost,
		Function:              controllers.SaveFile,
		RequestAuthentication: false,
	},
	
}
