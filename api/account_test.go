package api

import (
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestApi_GetAccount(t *testing.T) {
	stub := httptest.NewServer(FakeResponse(200, "../fixture/accounts.json"))
	FakeApi.BaseUrl = stub.URL

	data, _ := FakeApi.GetAccount()
	assert.NotEmpty(t, data.Account.ID)
}
