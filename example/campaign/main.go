package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func campaigns(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
		Graph:   1,
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	act := c.Param("act")
	url := fmt.Sprintf("%s/campaigns", act)
	resp := fbSDK.Get(url)
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func insights(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
		Graph:   1,
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	id := c.Param("id")
	url := fmt.Sprintf("%s/insights", id)
	resp := fbSDK.Get(url, sdk.WithParamQuery(sdk.ParamQuery{
		"fields": "reach",
	}))
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/campaigns/:act", campaigns)
	e.GET("/insights/:id", insights)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.Start(":8989"))
}
