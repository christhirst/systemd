package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func routes(e *echo.Echo, app *application) *echo.Echo {

	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get")
	})
	e.GET("/user", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users/get")
	})

	e.GET("/test", app.sendData)

	e.Logger.Printf("dd")
	e.Start(":8888")
	return e
}
