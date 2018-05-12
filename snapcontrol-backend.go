package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func echoHelloHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
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
