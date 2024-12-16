package api

import (
	"encoding/json"
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

type StopRequest struct {
	ProductId int
	RegGbn    string
	ModGbn    string
	RegId     string
}

func (ga GsshopApi) Stop(data []byte) error {
	var req StopRequest
	err := json.Unmarshal(data, &req)
	if err != nil {
		return err
	}

	path := "/alia/aliaCommonPrd.gs"
	params := url.Values{
		"regGbn":             {req.RegGbn},
		"modGbn":             {req.ModGbn},
		"regId":              {req.RegId},
		"supPrdCd":           {},
		"supCd":              {},
		"saleEndDtm":         {},
		"attrSaleEndStModYn": {"N"},
	}

	resp, err := ga.client.PostForm(path, params)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}
