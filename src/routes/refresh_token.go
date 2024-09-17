package routes

import (
	"api/src/controllers"
	"net/http"
)

var refresh_token = []Routes{
	{
		URI:                   "/refresh_token/{user_id}",
		Method:                http.MethodGet,
		Function:              controllers.RefreshToken,
		RequestAuthentication: false,
	},
	{
		URI:                   "/refresh_token",
		Method:                http.MethodGet,
		Function:              controllers.Teste,
		RequestAuthentication: false,
	},
}
