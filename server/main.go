package main

import (
	"echo_mgo_rest_server/server/routes"
	"fmt"
	"github.com/labstack/echo"
)

const port string = ":8080"

func main() {
	fmt.Printf("Running at %v\n", port)

	e := echo.New()

	routes.Init(e)

	e.Logger.Fatal(e.Start(port))
}
