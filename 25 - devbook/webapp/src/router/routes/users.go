package routes

import (
	"net/http"
	"webapp/src/controllers"
)

var userRoutes = []Route{
	{
		URI:         "/create-user",
		Method:      http.MethodGet,
		Func:        controllers.LoadCreateUserPage,
		AuthRequire: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Func:        controllers.CreateUser,
		AuthRequire: false,
	},
}
