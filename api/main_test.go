package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

var confFile, _ = ioutil.ReadFile("../conf.sample")
var FakeApi = Api{}
var _ = json.Unmarshal(confFile, &FakeApi)

func FakeResponse(httpStatus int, filename string) http.HandlerFunc {
	stub := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(httpStatus)
		body, _ := ioutil.ReadFile(filename)
		w.Write(body)
	})

	return stub
}

func init() {
	var _ = json.Unmarshal(confFile, &FakeApi)
}