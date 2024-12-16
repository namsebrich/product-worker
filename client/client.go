package client

import (
	"product-worker/client/api"
)

type Client struct {
	apis map[string]MarketApi
}

func New() *Client {
	apis := map[string]MarketApi{
		"gsshop": api.Gsshop(),
		"coupang": api.Coupang(),
	}

	return &Client{apis}
}

func (c *Client) Api(market string) (api MarketApi, exists bool) {
	api, exists = c.apis[market]
	return
}