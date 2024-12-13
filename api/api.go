package api

type MarketApi interface {
	Stop(data any) error
}

type Client struct {
	apis map[string]MarketApi
}

func NewClient() *Client {
	apis := map[string]MarketApi{
		"gsshop": GsshopApi(),
		"coupang": CoupangApi(),
	}

	return &Client{apis}
}

func (c *Client) Api(market string) (api MarketApi, exists bool) {
	api, exists = c.apis[market]
	return
}