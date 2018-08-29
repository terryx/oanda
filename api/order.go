package api

import (
	"fmt"
	"github.com/terryx/oanda/datastruct"
)

func (a *Api) CreateOrder(form map[string]datastruct.OrderForm) (datastruct.OrderCreateResponse, error) {
	req := a.setupReq("POST")
	req.Body = form
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s/orders", a.BaseUrl, a.AccountID)

	res, err := req.Do()

	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.OrderCreateResponse{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (a *Api) CancelOrder(id string) (datastruct.OrderCancelResponse, error) {
	req := a.setupReq("PUT")
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s/orders/%s/cancel", a.BaseUrl, a.AccountID, id)

	res, err := req.Do()

	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.OrderCancelResponse{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
