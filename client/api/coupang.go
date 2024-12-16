package api

import (
	"net/http"
)

type CoupangApi struct {
	client *http.Client
}

func Coupang() *CoupangApi {
	return &CoupangApi{}
}

func (c CoupangApi) Stop(data []byte) error {
	return nil
}
