package sdk

import (
	"errors"
	"fmt"
)

const bearerKey = "Bearer"

// Graph is wrap data for graph
type Graph struct {
	url         string
	appID       string
	appKey      string
	redirectURL string
	token       string
	secretProof string
}

// NewGraph is create new graph object
func NewGraph(options ...GraphOptions) *Graph {
	var graph = &Graph{}
	for _, option := range options {
		option(graph)
	}

	return graph
}

// ChangeToken is change token
func (g *Graph) ChangeToken(token string) error {
	if token == "" {
		return errors.New("fbSDK: error facebook access token is empty")
	}

	g.token = fmt.Sprintf("%s %s", bearerKey, token)
	return nil
}

// GraphOptions is options
type GraphOptions func(*Graph)

// GWithURL is paramter url facebook
func GWithURL(url string) GraphOptions {
	return func(g *Graph) {
		g.url = url
	}
}

// GWithAppID is parameter app id
func GWithAppID(ID string) GraphOptions {
	return func(g *Graph) {
		g.appID = ID
	}
}

// GWithAppKey is parameter app key
func GWithAppKey(key string) GraphOptions {
	return func(g *Graph) {
		g.appKey = key
	}
}

// GWithRedirectURL is parameter token
func GWithRedirectURL(url string) GraphOptions {
	return func(g *Graph) {
		g.redirectURL = url
	}
}

// GWithToken is parameter token
func GWithToken(token string) GraphOptions {
	return func(g *Graph) {
		g.token = fmt.Sprintf("%s %s", bearerKey, token)
	}
}

func gSecretProof(key, token string) GraphOptions {
	return func(g *Graph) {
		g.secretProof = HashHmac(key, []byte(token))
	}
}
