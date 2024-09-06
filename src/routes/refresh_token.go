package routes

import (
	"api/src/controllers"
	"net/http"
)

var refresh_token = Routes{
	URI: "/refrest_token/{user_id}",
	Method: http.MethodGet,
	Function: controllers.RefreshToken,
	RequestAuthentication: false,
}