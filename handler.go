package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/labstack/echo"
)

func (app *application) sendData(c echo.Context) error {
	id := c.QueryParam("components")
	fmt.Println(id)
	ctx := context.Background()
	names := []string{
		"pulseaudio.service",
	}
	con, err := dbus.NewUserConnectionContext(ctx)
	res, _ := con.ListUnitsByNamesContext(ctx, names)
	fmt.Println(con)
	if err != nil {
		fmt.Println(err)
	}
	//cmd.Execute()
	defer con.Close()
	return c.JSON(http.StatusOK, res)
}
