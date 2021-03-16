package core

import (
	"net/url"
	"os"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (c Config) Validate() (err error) {
	if err = is.URL.Validate(c.DiscoveryURL.String()); err != nil {
		return err
	}

	return validation.ValidateStruct(&c,
		validation.Field(&c.ClientID, validation.Required),
		validation.Field(&c.ClientSecret, validation.Required),
	)
}

func LoadConfig() (cfg Config) {
	if rawURL := os.Getenv("DISCOVERY_URL"); rawURL != "" {
		discoveryURL, _ := url.Parse(rawURL)

		cfg.DiscoveryURL = *discoveryURL
	}

	cfg.ClientID = os.Getenv("CLIENT_ID")
	cfg.ClientSecret = os.Getenv("CLIENT_SECRET")

	return cfg
}
