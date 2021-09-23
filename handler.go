package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"

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

func process(wg *sync.WaitGroup, u *action, con *dbus.Conn, ctx context.Context, ch chan<- string) {
	defer wg.Done()
	for _, v := range u.SERVICE {

		switch u.DO {
		case "restart":
			con.RestartUnitContext(ctx, v, "replace", ch)
			fmt.Println("Today.")
		case "stop":
			con.StopUnitContext(ctx, v, "replace", ch)
			fmt.Println("Tomorrow.")
		case "start":
			con.StartUnitContext(ctx, v, "replace", ch)
			fmt.Println("In two days.")
		case "kill":
			con.KillUnitContext(ctx, v, 9)
		default:
			close(ch)
		}
	}

}

func (app *application) sendAction(c echo.Context) error {
	var wg = sync.WaitGroup{}

	u := &action{}

	if err := c.Bind(u); err != nil {
		return err
	}

	wg.Add(len(u.SERVICE))
	ctx := context.Background()
	ch := make(chan string)
	con, err := dbus.NewUserConnectionContext(ctx)
	go process(&wg, u, con, ctx, ch)
	wg.Wait()

	//res, _ := con.ListUnitsByNamesContext(ctx, u.SERVICE)

	fmt.Println(u)

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
