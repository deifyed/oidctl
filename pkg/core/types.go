package core

import "net/url"

type Config struct {
	DiscoveryURL url.URL

	ClientID     string
	ClientSecret string
}

type DiscoveryDocument struct {
	Issuer        string `json:"issuer"`
	TokenEndpoint string `json:"token_endpoint"`
}

type AccessToken struct {
	Data      string `json:"access_token"`
	ExpiresIn int    `json:"expires_in"`
	Type      string `json:"token_type"`
}

type Flow interface {
	Authenticate(discoveryDocument DiscoveryDocument, clientID, clientSecret string) (token AccessToken, err error)
}
