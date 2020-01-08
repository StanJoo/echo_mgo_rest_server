package routes

import (
	todoroute "echo_mgo_rest_server/server/api/todo/route"
	userroute "echo_mgo_rest_server/server/api/user/route"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init(e *echo.Echo) {
	e.Pre(middleware.RemoveTrailingSlash())

	todoroute.Init(e)
	userroute.Init(e)
}
