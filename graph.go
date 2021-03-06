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
	graphType   int // node, edge
}

// NewGraph is create new graph object
func NewGraph(options ...GraphOptions) *Graph {
	var graph = &Graph{}
	for _, option := range options {
		option(graph)
	}

	return graph
}

// ChangeGraph is change graph type, available type is node or edge
func (g *Graph) ChangeGraph(graph int) {
	if graph < 0 || graph > 1 {
		return
	}

	g.graphType = graph
}

// ChangeToken is change token
func (g *Graph) ChangeToken(token string) error {
	if token == "" {
		return errors.New("fbSDK: error facebook access token is empty")
	}

	g.token = fmt.Sprintf("%s %s", bearerKey, token)
	g.changeSecretProof(token)
	return nil
}

// changeSecretProof is change secret proof
func (g *Graph) changeSecretProof(token string) {
	g.secretProof = HashHmac(g.appKey, []byte(token))
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

// GWithGraph is parameter token
func GWithGraph(graph int) GraphOptions {
	return func(g *Graph) {
		g.graphType = graph
	}
}

func gSecretProof(key, token string) GraphOptions {
	return func(g *Graph) {
		g.secretProof = HashHmac(key, []byte(token))
	}
}
