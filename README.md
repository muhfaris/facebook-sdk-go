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

### Both
Declaration field for Nodes and Edges in SDK is "Graph". Default value if You not declaration request type is `0`, So if You request Edges API set Graph to `1`.

Default value of "Graph" is:
 - 0 for Nodes
 - 1 for Edges

If you want use Edges API, the SDK config like:
```
config := sdk.Facebook{
	Token:   "<you facebook access token>",
	AppKey:  "<your app secret>",
	Version: "v9.0",
	Graph:   1,
}

fbSDK, err := sdk.NewFacebook(config)
if err != nil {
	// TODO error
}

resp := fbSDK.Get("<campaign_id>/insights", sdk.WithParamQuery(sdk.ParamQuery{"fields": "reach"}))
```

### APP Secret Proof
When you enable the `appsecret_proof` in Your app's settings, `AppKey` is must fill. You can not empty this field.

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

### GET Operation
This example request list ad account from `/me/adaccounts`.

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

### Post Operation
Post operation can use `WithBody()` for pass of data.

```
data := map[string]interface{}{
	"name":                  "test-1",
	"special_ad_categories": []string{"NONE"},
	"objective":             "CONVERSIONS",
	"status":                "PAUSED",
}

resp := got.Post("<act_id>/campaigns", WithBody(data))
if resp .Error != nil {
    // TODO error
	return
}

campaign := struct {
	ID string `json:"id,omitempty"`
}{}

_ = resp.Unmarshal(&campaign)
fmt.Println("Campaign ID:",  campaign.ID)
```

### Delete Operation
```
resp := got.Delete(dc.ID)
campaignResponse := struct {
	Success bool `json:"success,omitempty"`
}{}

_ = resp.Unmarshal(&campaignResponse)
fmt.Println("Response:", campaignResponse.Success)
```

### Batch Request
```
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

resp := got.Batch(batch)

if _, ok := resp.Data.HasErrors(); ok {
	// Todo Error
	return
}
```

## Contributing
Feel free to create an issue or send me a pull request if you have any "how-to" question or bug or suggestion when using this package. I'll try my best to reply to it.

## License
This package is licensed under the [MIT license](https://choosealicense.com/licenses/mit/).
