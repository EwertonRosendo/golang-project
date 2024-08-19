package routes

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Routes{
	{
		URI:                   "/users",
		Method:                http.MethodPost,
		Function:              controllers.CreateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users",
		Method:                http.MethodGet,
		Function:              controllers.FindUsers,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{user_id}",
		Method:                http.MethodGet,
		Function:              controllers.FindUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{user_id}",
		Method:                http.MethodPut,
		Function:              controllers.UpdateUser,
		RequestAuthentication: false,
	},
	{
		URI:                   "/users/{user_id}",
		Method:                http.MethodDelete,
		Function:              controllers.DeleteUser,
		RequestAuthentication: false,
	},
}
