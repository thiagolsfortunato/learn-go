package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Route{
	{
		URI:         "/posts",
		Method:      http.MethodPost,
		Func:        controllers.CreatePost,
		AuthRequire: true,
	},
	{
		URI:         "/posts",
		Method:      http.MethodGet,
		Func:        controllers.GetPosts,
		AuthRequire: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodGet,
		Func:        controllers.GetPost,
		AuthRequire: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodPut,
		Func:        controllers.UpdatePost,
		AuthRequire: true,
	},
	{
		URI:         "/posts/{postId}",
		Method:      http.MethodDelete,
		Func:        controllers.DeletePost,
		AuthRequire: true,
	},
	{
		URI:         "/users/{userId}/posts",
		Method:      http.MethodGet,
		Func:        controllers.GetPostByUser,
		AuthRequire: true,
	},
	{
		URI:         "/posts/{postId}/like",
		Method:      http.MethodPost,
		Func:        controllers.LikePost,
		AuthRequire: true,
	},
	{
		URI:         "/posts/{postId}/dislike",
		Method:      http.MethodPost,
		Func:        controllers.DislikePost,
		AuthRequire: true,
	},
}
