package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/terryx/oanda/datastruct"
	"net/http/httptest"
	"testing"
)

func TestApi_GetPricing(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/pricing.json"))
	FakeApi.BaseUrl = stub.URL

	instruments := "NAS100_USD"
	data, _ := FakeApi.GetPricing(datastruct.PricingParam{
		Instruments: instruments,
	})
	assert.NotEmpty(t, data.Time)
	assert.Equal(t, data.Prices[0].Instrument, instruments)
}
