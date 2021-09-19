package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/gitpod/mycli/cmd"
	coreDbus "github.com/godbus/dbus/v5"
	"github.com/labstack/echo"
)

func (app *application) sendData(c echo.Context) error {
	cc, _ := coreDbus.Connect("localhost:8888")
	fmt.Println("cc")
	fmt.Println(cc)

	ctx := context.Background()
	names := []string{
		"pulseaudio.service",
	}
	con, err := dbus.NewUserConnectionContext(ctx)
	res, _ := con.ListUnitsByNamesContext(ctx, names)
	fmt.Println()
	fmt.Println(con)
	fmt.Println(err)
	cmd.Execute()
	defer con.Close()

	fmt.Println("sss")
	return c.JSON(http.StatusOK, res)
}
