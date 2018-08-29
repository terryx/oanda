package api

import (
	"github.com/stretchr/testify/assert"
	"github.com/terryx/oanda/datastruct"
	"github.com/terryx/oanda/jsonparser"
	"log"
	"net/http/httptest"
	"testing"
)

func TestApi_GetInstruments(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/instruments.json"))
	FakeApi.BaseUrl = stub.URL

	data, err := FakeApi.GetInstruments()
	if err != nil {
		log.Fatal(err)
	}

	assert.NotNil(t, data)
}

func TestApi_GetCandles(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/candles.json"))
	FakeApi.BaseUrl = stub.URL

	query := datastruct.CandleQueryParam{
		Granularity: "M15",
		Count:       5000,
	}
	data, err := FakeApi.GetCandles("NAS100_USD", query)
	if err != nil {
		log.Fatal(err)
	}

	jsonparser.PrettyPrint(data.Candles)

	assert.NotNil(t, data)
}
