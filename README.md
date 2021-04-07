[![Go Report Card](https://goreportcard.com/badge/github.com/muhfaris/facebook-sdk-go)](https://goreportcard.com/report/github.com/muhfaris/facebook-sdk-go)

## Facebook SDK Go
The Facebook SDK Go is a library with powerful feature that enable Go developer to easily integrate Facebook login and make requests to the Graph API.

## Feature
- Authentication
- Graph API
- Marketing API
- Batch Request

## Installation
### Special for Login Authentication
- Please Setup HTTPS in your local machine
    - Locally trusted development tool (e.g mkcert)
    *Referenced:* [Requiring HTTPS for Facebook Login](https://developers.facebook.com/blog/post/2018/06/08/enforce-https-facebook-login/)

## Usage
---
### Overview
The Graph API is named after the idea of a "social graph" — a representation of the information on Facebook. It's composed of:
 - nodes — basically individual objects, such as a User, a Photo, a Page, or a Comment
 - edges — connections between a collection of objects and a single object, such as Photos on a Page or Comments on a Photo

#### Nodes
Reading operations almost always begin with a node. A node is an individual object with a unique ID. For example, there are many User node objects, each with a unique ID representing a person on Facebook. To read a node, you query a specific object's ID. So, to read your User node you would query its ID:
```
curl -i -X GET "https://graph.facebook.com/{your-user-id}?fields=id,name&access_token={your-user-access-token}"
```
This request would return the following fields (node properties) by default, formatted using JSON:

```
{
  "name": "Your Name",
  "id": "your-user-id"
}
```

#### Edges
Nodes have edges, which usually can return collections of other nodes which are attached to them. To read an edge, you must include both the node ID and the edge name in the path. For example, /user nodes have a /feed edge which can return all Post nodes on a User. You'll need to get a new access token and select user_posts permissions during the Get access token flow. Here's how you could use the edge to get all your Posts:
```
curl -i -X GET "https://graph.facebook.com/{your-user-id}/feed?access_token={your-user-access-token}"
```

The JSON response would look something like this:
```
{
  "data": [
    {
      "created_time": "2017-12-08T01:08:57+0000",
      "message": "Love this puzzle. One of my favorite puzzles",
      "id": "post-id"
    },
    {
      "created_time": "2017-12-07T20:06:14+0000",
      "message": "You need to add grape as a flavor.",
      "id": "post-id"
    }
  ]
}
```

### Both (Nodes and Edge)
Facebook SDK declaration Nodes and Edges with value 0 and 1:
 - 0 for Nodes
 - 1 for Edges

 Default value of SDK is Nodes (1), If you want to request Edges change `Graph` to 1. the SDK config like below:

```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func insights(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
		Graph: 1,
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := fbSDK.Get("<campaign_id>/insights", sdk.WithParamQuery(sdk.ParamQuery{"fields": "reach"}))
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func main(){
 	e := echo.New()

	// Routes
	e.GET("/insights", insights)
	e.Logger.Fatal(e.Start(":8989"))
}


```

### APP Secret Proof
When you enable the `appsecret_proof` in Your app's settings, `AppKey` is must fill. You can not empty this field.

### Login Authentication
Create two API, first use for request to facebook and second for retrive data from facebook.
```
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
```
*Note: Please, generate tls certificate your self*.

### GET Operation
This example request list ad account from `/me/adaccounts`.

```
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
```

if you want to request with some fields, just add param query in request:
```
resp := fbSDK.Get("/me/adaccounts", WithParamQuery(ParamQuery{
	"fields": "name,account_status",
}))
```

### Post Operation
Post operation can use `WithBody()` for pass of data.

```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func campaign(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	data := map[string]interface{}{
    	"name":                  "test-1",
    	"special_ad_categories": []string{"NONE"},
    	"objective":             "CONVERSIONS",
    	"status":                "PAUSED",
    	}

	resp := fbSDK.Post("<act_id>/campaigns", WithBody(data))
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	// Routes
	e.GET("/campaigns", campaign)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.Start(":8989"))
}

```

### Delete Operation
```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func campaigns(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	resp := fbSDK.Delete(<object_id>)
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	// Routes
	e.DELETE("/campaigns", campaigns)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.Start(":8989"))
}

```

### Batch Request
```
package main

import (
	"net/http"

	"github.com/labstack/echo/v4"

	sdk "github.com/muhfaris/facebook-sdk-go"
)

func batch(c echo.Context) error {
	config := sdk.Facebook{
		Token:   "",
		AppKey:  "",
		Version: "v9.0",
	}

	fbSDK, err := sdk.NewFacebook(config)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	batch := ArrayOfBatchBodyRequest{
    	{
    		Method:      "GET",
    		RelativeURL: "/me/adaccountssx",
    	},
    	{
    		Method:      "GET",
    		RelativeURL: "<act_id>/campaigns",
    	},
    }

	resp := fbSDK.Batch(batch)
	if resp.Error != nil {
		return c.JSON(http.StatusBadRequest, resp)
	}

	if _, ok := resp.Data.HasErrors(); ok {
	    // Todo Error
    }

	return c.JSON(http.StatusOK, resp)
}

func main() {
	e := echo.New()

	// Routes
	e.DELETE("/batch", batch)

	// Start server
	// generate ssl certificate use mkcert (https://github.com/FiloSottile/mkcert)
	e.Logger.Fatal(e.Start(":8989"))
}

```

## Contributing
Feel free to create an issue or send me a pull request if you have any "how-to" question or bug or suggestion when using this package. I'll try my best to reply to it.

## License
This package is licensed under the [MIT license](https://choosealicense.com/licenses/mit/).
