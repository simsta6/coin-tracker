package coinlore

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	timeout = time.Second * 30
)

//Client is
type Client struct {
	BaseURL    string
	HTTPClient *http.Client
}

//Coinlorer is
type Coinlorer interface {
	GetCurrency(ctx context.Context, cryptoID string) (currency Currency, err error)
}

//NewClient is
func NewClient(baseURL string) *Client {
	return &Client{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: timeout,
		},
	}
}

//GetCurrency is
func (c *Client) GetCurrency(ctx context.Context, cryptoID string) (currency RawCurrency, err error) {
	var rawCurrency []RawCurrency
	finalURL := fmt.Sprintf("%s/ticker/?id=%s", c.BaseURL, cryptoID)

	req, err := http.NewRequest("GET", finalURL, nil)
	if err != nil {
		return RawCurrency{}, err
	}

	req = req.WithContext(ctx)

	if err := c.sendRequest(req, &rawCurrency); err != nil {
		return RawCurrency{}, err
	}

	return rawCurrency[0], err
}

func (c *Client) sendRequest(req *http.Request, v interface{}) (err error) {
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Accept", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("Request %s: received %d status code", req.URL, res.StatusCode)
	}

	if err = json.NewDecoder(res.Body).Decode(&v); err != nil {
		return err
	}

	return err
}
