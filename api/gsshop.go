package api

import (
	"log"
	"net/http"
	"net/url"
)

type Gsshop struct {
	client *http.Client
}

func GsshopApi() *Gsshop {
	proxyUrl, err := url.Parse("http://172.31.17.167:3128")

	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		},
	}

	return &Gsshop{client}
}

func (gs Gsshop) Stop(data any) error {
	params := url.Values{
		"regGbn": {"U"},
		"modGbn": {"S"},
		"regId": {"BRT"},
		"supPrdCd": {},
		"supCd": {},
		"saleEndDtm": {},
		"attrSaleEndStModYn": {"N"},
	}

	resp, err := gs.client.PostForm("url", params)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}