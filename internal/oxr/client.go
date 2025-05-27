package oxr

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	APIHost string
	AppID   string
}

func NewClient(apiHost, appID string) *Client {
	return &Client{
		APIHost: apiHost,
		AppID:   appID,
	}
}

func (c Client) GetRates(currencies []string) (RatesResponse, error) {
	encAppID := url.PathEscape(c.AppID)
	symbols := strings.Join(currencies, ",")
	path := fmt.Sprintf("/latest.json?app_id=%s&base=USD&symbols=%s", encAppID, symbols)

	res, err := c.doRequest(http.MethodGet, path, nil)
	if err != nil {
		return RatesResponse{}, fmt.Errorf("getting rates: %w", err)
	}

	r := RatesResponse{}

	err = json.Unmarshal(res, &r)
	if err != nil {
		return RatesResponse{}, err
	}

	return r, nil
}

func (c Client) doRequest(method string, path string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, c.APIHost+path, bytes.NewBuffer(data))
	if err != nil {
		slog.Error("Failed to create HTTP request", "err", err)
		return nil, err
	}

	slog.Debug("Invoking Open Exchange Rates API", "URL", req.URL.String())

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		slog.Error("Failed to make HTTP request", "err", err)
		return nil, err
	}

	slog.Debug("Got response", "status code", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		slog.Error("Failed to read response body", "err", err)
		return nil, err
	}

	slog.Debug("Open Exchange Rates response", "body", string(resBody))

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return nil, fmt.Errorf("unexpected status code: %d", res.StatusCode)
	}

	return resBody, nil
}
