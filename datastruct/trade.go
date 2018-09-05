package datastruct

type TradesClosed struct {
	TradeID                string `json:"tradeID"`
	Units                  string `json:"units"`
	RealizedPL             string `json:"realizedPL"`
	Financing              string `json:"financing"`
	Price                  string `json:"price"`
	GuaranteedExecutionFee string `json:"guaranteedExecutionFee"`
	HalfSpreadCost         string `json:"halfSpreadCost"`
}

type TradeClosedResponse struct {
	OrderFillTransaction struct {
		TradesClosed []TradesClosed `json:"tradesClosed"`
	} `json:"orderFillTransaction"`
}
