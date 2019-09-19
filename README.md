# Unofficial Facebook SDK for Go
The Facebook SDK for Go is a library with powerful feature that enable
Go developer to easily integrate Facebook login and make requests
to the Graph API.

## Installation
- *Recommend* [Requiring HTTPS for Facebook Login](https://developers.facebook.com/blog/post/2018/06/08/enforce-https-facebook-login/)
- Locally trusted development tool (e.g mkcert)

## Usage
```
package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/muhfaris/facebook-sdk-go"
)

var (
	fbApp *facebook.App
)

func main() {
	fbApp = facebook.Init(
		"",
		"<application id>",
		"<application key>",
		"<callback url>",
		"<scopes>")

	r := mux.NewRouter()
	r.HandleFunc("/facebook", handlerFacebook).Methods(http.MethodGet)
	r.HandleFunc("/callback", handlerCallback).Methods(http.MethodGet)

	log.Println(http.ListenAndServe(":7070", r))
}

func handlerFacebook(w http.ResponseWriter, r *http.Request) {
	url := fbApp.GetAuthURL()
	http.Redirect(w, r, url, http.StatusFound)
	return
}

func handlerCallback(w http.ResponseWriter, r *http.Request) {
	code := r.FormValue("code")

	token, err := fbApp.GetToken(code)
	if err != nil {
		w.Write([]byte("Invalid to get token facebook"))
		return
	}

	fbApp.SetToken(token.AccessToken)
	data, err := fbApp.GET(
		"me",
		facebook.Params{
			"fields": "first_name, middle_name, last_name, name, email, short_name",
		},
	)

	response, err := json.Marshal(data)
	if err != nil {
		w.Write([]byte("Error can not parse response facebook"))
		return
	}

	w.Write(response)
	return
}
```


Response :
```
{
    email: "akunsosmedx02@gmail.com",
    first_name: "Muhammad",
    last_name: "Faris",
    name: "Muhammad Faris",
    short_name: "Muhammad"
}
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)
