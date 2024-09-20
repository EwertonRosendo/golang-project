package routes

import (
	"api/src/controllers"
	"net/http"
)

var reviews = []Routes{
	{
		URI:                   "/reviews/{review_id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteReview,
		RequestAuthentication: true,
	},
	{
		URI:                   "/reviews",
		Method:                http.MethodGet,
		Function:              controllers.SearchReviews,
		RequestAuthentication: false,
	},
	{
		URI:                   "/reviews/{review_id}",
		Method:                http.MethodGet,
		Function:              controllers.FindReviewById,
		RequestAuthentication: false,
	},
	{
		URI:                   "/reviews/users/{user_id}",
		Method:                http.MethodGet,
		Function:              controllers.FindReviewsByUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/reviews/add",
		Method:                http.MethodPost,
		Function:              controllers.AddReview,
		RequestAuthentication: true,
	},
}
