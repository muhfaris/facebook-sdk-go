package facebook

import (
	"fmt"
	"net/http"

	"golang.org/x/oauth2"
)

type App struct {
	name        string
	appID       string
	appKey      string
	authURL     string
	tokenURL    string
	callbackURL string
	scopes      []string
	token       string
	config      *oauth2.Config
	httpClient  http.Client
	version     string
}

// Init is initialize init function
func Init(name, appID, appKey, callbackURL, version string, scopes ...string) *App {
	a := &App{
		name:        name,
		appID:       appID,
		appKey:      appKey,
		callbackURL: callbackURL,
		version:     version,
	}

	a.config = newConfig(a, scopes)
	return a
}

// InitRequest if use without auth facebook
func InitRequest(name, appID, appKey, version, token string) *App {
	a := &App{
		name:    name,
		appID:   appID,
		appKey:  appKey,
		version: version,
		token:   token,
	}

	return a
}

func newConfig(app *App, scopes []string) *oauth2.Config {
	c := &oauth2.Config{
		ClientID:     app.appID,
		ClientSecret: app.appKey,
		RedirectURL:  app.callbackURL,
		Endpoint: oauth2.Endpoint{
			AuthURL:  DefaultauthURL,
			TokenURL: DefaulttokenURL,
		},
	}

	for _, scope := range scopes {
		if _, exists := defaultScopes[scope]; !exists {
			c.Scopes = append(c.Scopes, scope)
		}
	}

	return c
}

func (app *App) Name() string {
	return app.name
}

func (app *App) GetAuthURL() string {
	return app.config.AuthCodeURL("state", oauth2.AccessTypeOnline)
}

func (app *App) GetToken(code string) (*oauth2.Token, error) {
	token, err := app.config.Exchange(oauth2.NoContext, code)
	if err != nil {
		return nil, err
	}

	app.SetToken(token.AccessToken)
	return token, nil
}

func (app *App) SetToken(token string) {
	app.token = fmt.Sprintf("%s %s", labelBearer, token)
}

func (app *App) getURIFacebook() string {
	var version = DefaultVersionAPI
	if app.version != "" {
		version = app.version
	}

	return fmt.Sprintf("%s/%s", FacebookAPI, version)
}
