package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func adaccount(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := fbSDK.Get("/me/adaccounts")
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/adaccounts", adaccount)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.Start(":8989"))
}
