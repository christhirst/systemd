package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/coreos/go-systemd/v22/dbus"
	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

func (app *application) sendData(c echo.Context) error {
	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://accounts.google.com")
	if err != nil {
		// handle error
	}

	// Configure an OpenID Connect aware OAuth2 client.
	oauth2Config := oauth2.Config{
		ClientID:     "clientID",
		ClientSecret: "clientSecret",
		RedirectURL:  "redirectURL",

		// Discovery returns the OAuth2 endpoints.
		Endpoint: provider.Endpoint(),

		// "openid" is a required scope for OpenID Connect flows.
		Scopes: []string{oidc.ScopeOpenID, "profile", "email"},
	}
	fmt.Println(oauth2Config)
	params := c.QueryParams()

	var stringL []string
	for _, v := range params {
		stringL = append(stringL, v[0])

	}
	fmt.Println(stringL)
	con, err := dbus.NewUserConnectionContext(ctx)
	res, _ := con.ListUnitsByNamesContext(ctx, stringL)
	con.RestartUnitContext(ctx, "gnome-keyring-ssh.service", "replace", nil)
	/* 	con.StopUnitContext()
	   	con.StartUnitContext()
	   	con.KillUnitContext() */

	fmt.Println(con)
	if err != nil {
		fmt.Println(err)
	}
	//cmd.Execute()
	defer con.Close()
	return c.JSON(http.StatusOK, res)
}
