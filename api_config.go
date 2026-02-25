package main

import (
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/CodeZeroSugar/ip-guard/internal/client"
)

const (
	baseURL = "https://api.abuseipdb.com/api/v2"
)

type config struct {
	abusedClient client.Client
	fullURL      string
	apiKey       string
}

func newConfig(key, endpoint string, params map[string]string) (*config, error) {
	if len(params) == 0 {
		return nil, errors.New("new config requires url parameters")
	}
	if endpoint != "check" {
		return nil, errors.New("currently only 'check' endpoint accepted")
	}

	p := url.Values{}
	for k, v := range params {
		p.Add(k, v)
	}

	endpointURL := baseURL + "/" + endpoint
	fullURL := fmt.Sprintf("%s?%s", endpointURL, p.Encode())

	return &config{
		abusedClient: client.NewClient(5 * time.Minute),
		fullURL:      fullURL,
		apiKey:       key,
	}, nil
}
