package main

import (
	"net/http"

	"github.com/labstack/echo"
	"io/ioutil"
)

func echoHelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		body, _ := ioutil.ReadAll(c.Request().Body)

		if len(body) > 0 {
			return c.String(http.StatusOK, "Hello World, Ulsmar!")
		}
		return c.String(http.StatusOK, "Hello World")
	}
}

func setupEchoEngine() *echo.Echo {
	e := echo.New()
	e.GET("/", echoHelloHandler())

	return e
}

func main() {
	e := setupEchoEngine()
	e.Logger.Fatal(e.Start(":1323"))
}
