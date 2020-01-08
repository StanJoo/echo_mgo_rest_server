package postroute

import (
	postcontroller "echo_mgo_rest_server/server/api/post/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("api/posts", postcontroller.GetAll)
	e.GET("api/posts/:title", postcontroller.GetByTitle)
	e.POST("api/posts", postcontroller.NewPost)
	e.DELETE("api/posts/:id", postcontroller.RemovePost)
}
