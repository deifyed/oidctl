package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func GetDiscoveryDocument(discoveryURL url.URL) (discoveryDocument DiscoveryDocument, err error) {
	var (
		response *http.Response
		payload  []byte
	)

	response, err = http.Get(discoveryURL.String())
	if err != nil {
		return discoveryDocument, fmt.Errorf("getting discovery document: %w", err)
	}

	defer func() {
		_ = response.Body.Close()
	}()

	payload, err = io.ReadAll(response.Body)
	if err != nil {
		return discoveryDocument, fmt.Errorf("reading body: %w", err)
	}

	err = json.Unmarshal(payload, &discoveryDocument)
	if err != nil {
		return discoveryDocument, fmt.Errorf("unmarshalling discovery document: %w", err)
	}

	return discoveryDocument, nil
}
