package datastruct

type Price struct {
	Type string `json:"type"`
	Time string `json:"time"`
	CloseoutBid string `json:"closeoutBid"`
	CloseoutAsk string `json:"closeoutAsk"`
	Status string `json:"status"`
	Tradeable bool `json:"tradeable"`
	Instrument string `json:"instrument"`
}

type Prices struct {
	Price []Price `json:"prices"`
}

type PricingParam struct {
	Instruments string `json:"instruments"`
}

type PricingResponse struct {
	Time string `json:"time"`
	Prices []Price `json:"prices"`
}