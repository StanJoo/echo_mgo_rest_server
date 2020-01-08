package route

import (
	usercontroller "echo_mgo_rest_server/server/api/user/controller"
	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	e.GET("api/users", usercontroller.GetAll)
	e.GET("api/users/:id", usercontroller.GetById)
	e.POST("api/users", usercontroller.NewUser)
	e.DELETE("api/users/:id", usercontroller.RemoveUser)
}
