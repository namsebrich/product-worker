package api

import (
	"net/http"
)

type Coupang struct {
	client *http.Client
}

func CoupangApi() *Coupang {
	return &Coupang{}
}

func (c Coupang) Stop(data any) error {
	return nil
}