package api

import (
	"fmt"
	"github.com/terryx/oanda/datastruct"
)

func (a *Api) GetAccount() (datastruct.AccountDetail, error) {
	req := a.setupReq("GET")
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s", a.BaseUrl, a.AccountID)

	res, err := req.Do()

	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.AccountDetail{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
