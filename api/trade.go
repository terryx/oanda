package api

import (
	"fmt"
	"github.com/terryx/oanda/datastruct"
)

func (a *Api) CloseTrade(tradeID string) (datastruct.TradeClosedResponse, error) {
	req := a.setupReq("PUT")
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s/trades/%s/close", a.BaseUrl, a.AccountID, tradeID)

	res, err := req.Do()

	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.TradeClosedResponse{}

	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
