package api

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestApi_CloseTrade(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/trade-closed.json"))
	FakeApi.BaseUrl = stub.URL

	data, _ := FakeApi.CloseTrade(string(4119))
	assert.Equal(t, data.OrderFillTransaction.TradesClosed[0].TradeID, "4119")
}
