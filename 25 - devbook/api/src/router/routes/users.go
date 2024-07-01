package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Route{
	{
		URI:         "/users",
		Method:      http.MethodPost,
		Func:        controllers.CreateUser,
		AuthRequire: false,
	},
	{
		URI:         "/users",
		Method:      http.MethodGet,
		Func:        controllers.GetUsers,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodGet,
		Func:        controllers.GetUser,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodPut,
		Func:        controllers.UpdateUser,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}",
		Method:      http.MethodDelete,
		Func:        controllers.DeleteUser,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/follow",
		Method:      http.MethodPost,
		Func:        controllers.FollowUser,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/unfollow",
		Method:      http.MethodPost,
		Func:        controllers.UnfollowUser,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/followers",
		Method:      http.MethodGet,
		Func:        controllers.GetFollowers,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/following",
		Method:      http.MethodGet,
		Func:        controllers.GetFollowing,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/password",
		Method:      http.MethodPost,
		Func:        controllers.UpdatePassword,
		AuthRequire: true,
	},
}
