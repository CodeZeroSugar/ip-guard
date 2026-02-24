// Package client contains logic for http client that queries AbusedIPDB
package client

import (
	"io"
	"log"
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{Timeout: timeout},
	}
}

func (c *Client) FetchIPData(url, key string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("%s", err)
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("Key", key)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Fatalf("%s", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return body, nil
}
