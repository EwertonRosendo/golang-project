package routes

import (
	"api/src/controllers"
	"net/http"
)

var books = []Routes{
	{
		URI:                   "/books/{book_id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteBook,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books",
		Method:                http.MethodGet,
		Function:              controllers.SearchBooks,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books/{book_id}",
		Method:                http.MethodGet,
		Function:              controllers.FindBookById,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books/{title}",
		Method:                http.MethodGet,
		Function:              controllers.SearchBooksByTitle,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books",
		Method:                http.MethodPost,
		Function:              controllers.AddBookWithFile,
		RequestAuthentication: false,
	},
	{
		URI:                   "/books/{book_id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateBookWithFile,
		RequestAuthentication: false,
	},
}
