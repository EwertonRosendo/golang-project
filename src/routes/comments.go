package routes

import (
	"api/src/controllers"
	"net/http"
)

var comments = []Routes{
	{
		URI:                   "/reviews/{review_id}/comments",
		Method:                http.MethodPost,
		Function:              controllers.CreateComment,
		RequestAuthentication: true,
	},
	{
		URI:                   "/reviews/{review_id}/comments",
		Method:                http.MethodGet,
		Function:              controllers.SearchComments,
		RequestAuthentication: false,
	},
	{
		URI:                   "/comments/{comment_id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteComment,
		RequestAuthentication: true,
	},
}
