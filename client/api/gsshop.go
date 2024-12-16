package api

import (
	"log"
	"net/http"
	"net/url"
)

type GsshopApi struct {
	client *http.Client
}

func Gsshop() *GsshopApi {
	proxyUrl, err := url.Parse("http://172.31.17.167:3128")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	return &GsshopApi{client}
}

func (ga GsshopApi) Stop(data any) error {
	params := url.Values{
		"regGbn": {"U"},
		"modGbn": {"S"},
		"regId": {"BRT"},
		"supPrdCd": {},
		"supCd": {},
		"saleEndDtm": {},
		"attrSaleEndStModYn": {"N"},
	}

	resp, err := ga.client.PostForm("url", params)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}