package client

type MarketApi interface {
	Stop(data []byte) error
}
