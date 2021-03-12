package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func adaccount(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "EAAI0jja7Wi0BAMg5dFZB9u1IxkQGPmL8S9IZB5SVZAhoBkmkTLPOWxYdaWJn3NVqcifCxOfeHsq4xzJ7OGwUyf3C6v5G4TZAmJgMg5QBchwLsp5PQBQOuXOVMPXHMWxszFio8EJYiqQdkf9NJcT31YTiWpXBCdFqW5w8rlZClWnx5RUvaiw4NQOZBthKB0UMO4ibPfdUccMgxaRz8a7VANtWoBOAHelGEZD",
		AppKey:  "30ba840a7c23cf8baea15669a09002c9",
		Version: "v9.0",
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := fbSDK.Get("/me/adaccounts", sdk.ParamQuery{})
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
