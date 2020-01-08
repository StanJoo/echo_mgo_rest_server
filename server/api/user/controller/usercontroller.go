package usercontroller

import (
	userdao "echo_mgo_rest_server/server/api/user/dao"
	user "echo_mgo_rest_server/server/api/user/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetAll(c echo.Context) error {
	us, _ := userdao.All()

	return c.JSON(http.StatusOK, us)
}

func GetById(c echo.Context) error {
	id := c.Param("id")

	nt, _ := userdao.GetById(id)

	return c.JSON(http.StatusOK, nt)
}

func NewUser(c echo.Context) error {
	u := new(user.User)

	//c.Bind(u)
	c.Bind(u)

	nt, _ := userdao.NewUser(*u)

	return c.JSON(http.StatusOK, nt)
}

func RemoveUser(c echo.Context) error {
	id := c.Param("id")

	userdao.DeleteUser(id)

	return c.String(http.StatusOK, "")
}
