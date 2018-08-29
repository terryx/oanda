package api

import (
	"fmt"
	"github.com/terryx/oanda/datastruct"
)

func (a *Api) GetInstruments() (datastruct.Instruments, error) {
	req := a.setupReq("GET")
	req.Uri = fmt.Sprintf("%s/v3/accounts/%s/instruments", a.BaseUrl, a.AccountID)

	res, err := req.Do()
	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.Instruments{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}

func (a *Api) GetCandles(instrument string, param datastruct.CandleQueryParam) (datastruct.Candles, error) {
	req := a.setupReq("GET")
	req.Uri = fmt.Sprintf("%s/v3/instruments/%s/candles", a.BaseUrl, instrument)
	req.QueryString = param

	res, err := req.Do()
	if res.StatusCode >= 400 {
		message, _ := res.Body.ToString()
		panic(message)
	}

	data := datastruct.Candles{}
	if err != nil {
		return data, err
	}

	err = res.Body.FromJsonTo(&data)

	if err != nil {
		return data, err
	}

	return data, nil
}
