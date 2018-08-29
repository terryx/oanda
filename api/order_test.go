package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/terryx/oanda/datastruct"
	"net/http/httptest"
	"testing"
)

func TestApi_CreateOrder(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/order-created.json"))
	FakeApi.BaseUrl = stub.URL

	param := map[string]datastruct.OrderForm{
		"order": {
			Type:         "LIMIT",
			Instrument:   "WHEAT_USD",
			Units:        10,
			Price:        "4.500",
			TimeInForce:  "GTC",
			PositionFill: "DEFAULT",
			StopLossOnFill: datastruct.StopLossOnFill{
				Price: "4.000",
			},
			TakeProfitOnFill: datastruct.TakeProfitOnFill{
				Price: "5.000",
			},
		},
	}

	data, _ := FakeApi.CreateOrder(param)
	assert.Equal(t, data.Transaction.Type, "LIMIT_ORDER")
}

func TestApi_CancelOrder(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/order-canceled.json"))
	FakeApi.BaseUrl = stub.URL

	data, _ := FakeApi.CancelOrder("30")
	assert.Equal(t, data.Transaction.OrderID, "30")
}
