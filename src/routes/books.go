package routes

import (
	"api/src/controllers"
	"net/http"
)

var books = []Routes{
	{
		URI:                   "/books",
		Method:                http.MethodGet,
		Function:              controllers.SearchBooks,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books/{title}",
		Method:                http.MethodGet,
		Function:              controllers.SearchBooksByTitle,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books/add",
		Method:                http.MethodPost,
		Function:              controllers.AddBook,
		RequestAuthentication: false,
	},
}
