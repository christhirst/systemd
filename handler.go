package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/labstack/echo"
)

type (
	action struct {
		ID      string   `json:"id"`
		DO      string   `json:"do"`
		SERVICE []string `json:"service"`
	}
)

func (app *application) sendAction(c echo.Context) error {

	u := &action{}

	if err := c.Bind(u); err != nil {
		return err
	}
	ctx := context.Background()

	con, err := dbus.NewUserConnectionContext(ctx)
	//res, _ := con.ListUnitsByNamesContext(ctx, u.SERVICE)

	ch := make(chan string)

	fmt.Println(u)
	switch u.DO {
	case "restart":
		con.RestartUnitContext(ctx, u.SERVICE[0], "replace", ch)
		fmt.Println("Today.")
	case "stop":
		con.StopUnitContext(ctx, u.SERVICE[0], "replace", ch)
		fmt.Println("Tomorrow.")
	case "start":
		con.StartUnitContext(ctx, u.SERVICE[0], "replace", ch)
		fmt.Println("In two days.")
	case "kill":
		con.KillUnitContext(ctx, u.SERVICE[0], 9)
	default:
		close(ch)
	}
	x := <-ch
	fmt.Println("eee")
	fmt.Println(x)
	if err != nil {
		fmt.Println(err)
	}
	//cmd.Execute()
	defer con.Close()
	return c.JSON(http.StatusOK, nil)
}
func (app *application) status(c echo.Context) error {

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
