package sdk

import (
	"golang.org/x/oauth2"
)

type Authenticate struct {
	oauth *oauth2.Config
	state string
}

// createAuthConfig is create new object of authentication
func createAuthConfig(appID, appKey, redirectURL, state string, scopes ...string) *Authenticate {
	return &Authenticate{
		oauth: &oauth2.Config{
			ClientID:     appID,
			ClientSecret: appKey,
			RedirectURL:  redirectURL,
			Scopes:       scopes,
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://www.facebook.com/dialog/oauth",
				TokenURL: "https://graph.facebook.com/oauth/access_token",
			},
		},
		state: state,
	}
}

// URL is authentication to facebook
func (a *Authenticate) URL() string {
	return a.oauth.AuthCodeURL(a.state)
}

// GetAccessToken is get access token facebook
func (a *Authenticate) GetAccessToken(code string) (*oauth2.Token, error) {
	return a.oauth.Exchange(oauth2.NoContext, code)
}
