# Unofficial Facebook SDK for Go
The Facebook SDK for Go is a library with powerful feature that enable
Go developer to easily integrate Facebook login and make requests
to the Graph API.

## Installation
- *Recommend* [Requiring HTTPS for Facebook Login](https://developers.facebook.com/blog/post/2018/06/08/enforce-https-facebook-login/)
- Locally trusted development tool (e.g mkcert)

## Usage
### Get Ad accounts (Graph API)
```
    import github.com/muhfaris/facebook-sdk-go"
    appID := "xxx"
	appSecret:= "yyyy"
	token := "token"

	fbSDK := facebook.InitRequest("test main", appID, appSecret, "v5.0", token)
	resp := fbSDK.Request(facebook.GraphAPI, "/me/adaccounts", http.MethodGet, nil)

	if resp.Error != nil {
		log.Println(resp.Error)
		return
	}

	log.Println("Response:", string(resp.Result.([]byte)))
```

### Get campaigns (Marketing API)
```
    import github.com/muhfaris/facebook-sdk-go"

    appID := "xxx"
	appSecret:= "yyy"
	token := "token"

	fbSDK := facebook.InitRequest("test main", appID, appSecret, "v5.0", token)
	resp := fbSDK.Request(
		facebook.GraphAPI,
		"/act_11111/campaigns",
		http.MethodGet,
		nil)

	if resp.Error != nil {
		log.Println(resp.Error)
		return
	}

	log.Println("Response:", string(resp.Result.([]byte)))
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
