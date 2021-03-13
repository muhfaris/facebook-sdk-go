## Facebook SDK GO
The Facebook SDK GO is a library with powerful feature that enable Go developer to easily integrate Facebook login and make requests to the Graph API.

## Feature
- Login Authentication
- Graph API
- Marketing API

## Installation
### Special for Login Authentication
- Please Setup HTTPS in your local machine
    - Locally trusted development tool (e.g mkcert)
    *Referenced:* [Requiring HTTPS for Facebook Login](https://developers.facebook.com/blog/post/2018/06/08/enforce-https-facebook-login/)

## Usage
---
### Login Authentication
Create two API, first use for request to facebook and second for retrive data from facebook.
```
var fbSDK sdk.FacebookAPI
// Handler
func login(c echo.Context) error {
	config := sdk.Facebook{
		AppID:       "",
		AppKey:      "",
		Version:     "v9",
		RedirectURL: "https://localhost:8989/callback", // reference to callback function
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
```

### List Ad Account
Create handler like below:
```
func adaccount(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "<your_facebook_access_token>",
		AppKey:  "<your_faebook_app_key/secret_app>",
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
```

if you want to request some fields, just add param query in request:
```
	resp := fbSDK.Get("/me/adaccounts", WithParamQuery(ParamQuery{
		"fields": "name,account_status",
	}))
```

## Contributing
Feel free to create an issue or send me a pull request if you have any "how-to" question or bug or suggestion when using this package. I'll try my best to reply to it.

## License
[MIT](https://choosealicense.com/licenses/mit/)
