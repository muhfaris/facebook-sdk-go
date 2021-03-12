package main

import (
	"net/http"

	sdk "github.com/muhfaris/facebook-sdk-go"

	"github.com/labstack/echo/v4"
)

var fbSDK sdk.FacebookAPI

// Handler
func login(c echo.Context) error {
	config := sdk.Facebook{
		AppID:       "",
		AppKey:      "",
		Version:     "v9",
		RedirectURL: "https://localhost:8989/callback",
	}

	fbSDK, _ = sdk.NewFacebook(config)

	authenticate := fbSDK.Authenticate("example-state")
	url := authenticate.URL()

	c.Redirect(http.StatusFound, url)
	return nil
}

func callback(c echo.Context) error {
	state := c.FormValue("state")
	authenticate := fbSDK.Authenticate(state)

	code := c.FormValue("code")
	token, _ := authenticate.GetAccessToken(code)

	return c.JSON(http.StatusOK, token)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/login", login)
	e.GET("/callback", callback)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.StartTLS(":8989", "localhost+2.pem", "localhost+2-key.pem"))
}
