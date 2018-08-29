package api

import (
	"fmt"
	"github.com/franela/goreq"
	"time"
)

type Api struct {
	ApiKey    string
	AccountID string
	BaseUrl   string
}

func (a *Api) setupReq(method string) goreq.Request {
	req := goreq.Request{
		Method: method,
	}
	req.Timeout = time.Duration(10 * time.Second)
	req.AddHeader("Content-Type", "application/json")
	req.AddHeader("Authorization", fmt.Sprintf("Bearer %s", a.ApiKey))

	return req
}
