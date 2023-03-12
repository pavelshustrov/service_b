package v1

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/pavelshustrov/service_b/pkg/client/v1/response"
)

type Client struct {
	client *resty.Client
}

func New() *Client {
	return &Client{
		client: resty.New(),
	}
}

// Create creates a new order.
func (client *Client) Create(
	ctx context.Context,
) (*response.ID, error) {
	resp, err := client.client.R().
		SetHeader("Content-Type", "application/json").
		Get("http://localhost:8777/uuid")

	if err != nil {
		return nil, err
	}

	var res response.ID

	if err := json.Unmarshal(resp.Body(), &res); err != nil {
		return nil, err
	}

	return &res, nil
}
