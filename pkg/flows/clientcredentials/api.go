package clientcredentials

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/deifyed/oidctl/pkg/core"
)

func NewClientCredentials(audience string) core.Flow {
	return &flow{audience: audience}
}

func (f flow) Authenticate(discoveryDocument core.DiscoveryDocument, clientID, clientSecret string) (token core.AccessToken, err error) {
	formData := url.Values{
		"grant_type":    {"client_credentials"},
		"client_id":     {clientID},
		"client_secret": {clientSecret},
		"audience":      {f.audience},
	}

	request, err := http.NewRequest(http.MethodPost, discoveryDocument.TokenEndpoint, strings.NewReader(formData.Encode()))
	if err != nil {
		return token, fmt.Errorf("creating request: %w", err)
	}

	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	client := http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return token, fmt.Errorf("doing request: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	rawBody, err := io.ReadAll(response.Body)
	if err != nil {
		return token, fmt.Errorf("reading body: %w", err)
	}

	err = json.Unmarshal(rawBody, &token)
	if err != nil {
		return token, fmt.Errorf("unmarshalling token: %w", err)
	}

	return token, nil
}
