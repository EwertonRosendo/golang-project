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
		RequestAuthentication: true,
	},
	{
		URI:                   "/books/{book_id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateBook,
		RequestAuthentication: true,
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
		URI:                   "/books/add",
		Method:                http.MethodPost,
		Function:              controllers.AddBook,
		RequestAuthentication: true,
	},
}
