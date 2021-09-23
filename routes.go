package main

import (
	"github.com/labstack/echo"
)

func routes(e *echo.Echo, app *application) *echo.Echo {

	e.POST("/status", app.status)

	e.POST("/manager", app.sendAction)

	e.Logger.Printf("dd")
	e.Start(":8888")
	return e
}
