package usercontroller

import (
	userdao "echo_mgo_rest_server/server/api/user/dao"
	user "echo_mgo_rest_server/server/api/user/model"
	"github.com/dgrijalva/jwt-go"
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
	//username := c.FormValue("username")

	userdao.DeleteUser(id)

	//userdao.DeleteUser(username)

	return c.String(http.StatusOK, "")
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	nt, _ := userdao.GetByUsername(username)

	// Throws Unauthorized Error
	if username != nt.Username || password != nt.Password {
		return echo.ErrUnauthorized
	}

	// Create Token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = nt.Username

	// Generate encoded token and send it as response
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
