package postcontroller

import (
	postdao "echo_mgo_rest_server/server/api/post/dao"
	post "echo_mgo_rest_server/server/api/post/model"
	"github.com/labstack/echo"
	"net/http"
)

func GetAll(c echo.Context) error {
	ps, _ := postdao.All()

	return c.JSON(http.StatusOK, ps)
}

func GetByTitle(c echo.Context) error {
	title := c.Param("title")

	nt, _ := postdao.GetByTitle(title)

	return c.JSON(http.StatusOK, nt)
}

func NewPost(c echo.Context) error {
	p := new(post.Post)

	c.Bind(p)

	nt, _ := postdao.NewPost(*p)

	return c.JSON(http.StatusOK, nt)
}

func RemovePost(c echo.Context) error {
	id := c.Param("id")

	postdao.DeletePost(id)

	return c.String(http.StatusOK, "")
}
