package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/labstack/echo/v4"
)

type (
	action struct {
		ID      string   `json:"id"`
		DO      string   `json:"do"`
		SERVICE []string `json:"service"`
	}

	response struct {
		status map[string]string
	}
)

func processWorker(ctx context.Context, wg *sync.WaitGroup, v string, do string, con *dbus.Conn, ch chan<- string) {
	defer wg.Done()
	switch do {
	case "restart":
		con.RestartUnitContext(ctx, v, "replace", ch)
	case "stop":
		con.StopUnitContext(ctx, v, "replace", ch)
		fmt.Println("Tomorrow.")
	case "start":
		con.StartUnitContext(ctx, v, "replace", ch)
		fmt.Println("In two days.")
	case "kill":
		con.KillUnitContext(ctx, v, 9)
	}

}

var wg = sync.WaitGroup{}

func (app *application) sendAction(c echo.Context) error {
	ctx := context.Background()

	u := &action{}
	if err := c.Bind(u); err != nil {
		return err
	}
	wg.Add(len(u.SERVICE))

	con, err := dbus.NewUserConnectionContext(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer con.Close()

	resp := new(response)
	resp.status = make(map[string]string)
	ch := make(map[int]chan string)
	for i, v := range u.SERVICE {
		ch[i] = make(chan string)
		go processWorker(ctx, &wg, v, u.DO, con, ch[i])
	}

	for i, _ := range u.SERVICE {
		v := <-ch[i]
		resp.status[u.SERVICE[i]] = v
	}
	wg.Wait()

	b, err := json.Marshal(resp.status)
	if err != nil {
		fmt.Println(err)
	}
	return c.JSON(http.StatusOK, string(b))
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
