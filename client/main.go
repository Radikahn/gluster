package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Radikahn/gluster/api/types"
)

type Client struct {
	BaseURL string
	HTTP    *http.Client
}

func (c *Client) CheckHealth(ctx context.Context) (*types.HealthResponse, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.BaseURL+"/api/system/health", nil)

	if err != nil {
		return nil, err
	}

	resp, err := c.HTTP.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("health check failed: %s", resp.Status)
	}

	var out types.HealthResponse
	return &out, json.NewDecoder(resp.Body).Decode(&out)
}

func main() {

}
