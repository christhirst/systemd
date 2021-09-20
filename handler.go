package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/labstack/echo"
)

func (app *application) sendData(c echo.Context) error {
	params := c.QueryParams()

	ctx := context.Background()
	var stringL []string
	for _, v := range params {
		stringL = append(stringL, v[0])

	}
	fmt.Println(stringL)
	con, err := dbus.NewUserConnectionContext(ctx)
	res, _ := con.ListUnitsByNamesContext(ctx, stringL)
	fmt.Println(con)
	if err != nil {
		fmt.Println(err)
	}
	//cmd.Execute()
	defer con.Close()
	return c.JSON(http.StatusOK, res)
}
