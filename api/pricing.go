package api

import (
	"fmt"
	"github.com/terryx/oanda/datastruct"
)

func (a *Api) GetPricing(param datastruct.PricingParam) (datastruct.PricingResponse, error) {
	req := a.setupReq("GET")
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s/pricing", a.BaseUrl, a.AccountID)
	req.QueryString = param

	res, err := req.Do()

	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.PricingResponse{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
