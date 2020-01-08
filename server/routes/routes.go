package routes

import (
	postroute "echo_mgo_rest_server/server/api/post/route"
	todoroute "echo_mgo_rest_server/server/api/todo/route"
	userroute "echo_mgo_rest_server/server/api/user/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// todo_route
	todoroute.Init(e)
	// user_route
	userroute.Init(e)
	//post_route
	postroute.Init(e)
}
