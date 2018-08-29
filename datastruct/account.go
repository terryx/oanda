package datastruct

type AccountOrder struct {
	ID               string `json:"id"`
	CreateTime       string `json:"createTime"`
	Type             string `json:"type"`
	TradeID          string `json:"tradeID"`
	Price            string `json:"price"`
	TimeInForce      string `json:"timeInForce"`
	TriggerCondition string `json:"triggerCondition"`
	State            string `json:"state"`
}

type AccountTrade struct {
	ID                    string `json:"id"`
	Instrument            string `json:"instrument"`
	Price                 string `json:"price"`
	OpenTime              string `json:"openTime"`
	InitialUnits          string `json:"initialUnits"`
	InitialMarginRequired string `json:"initialMarginRequired"`
	State                 string `json:"state"`
	CurrentUnits          string `json:"currentUnits"`
	RealizedPL            string `json:"realizedPL"`
	Financing             string `json:"financing"`
	TakeProfitOrderID     string `json:"takeProfitOrderID"`
	StopLossOrderID       string `json:"stopLossOrderID"`
	UnrealizedPL          string `json:"unrealizedPL"`
	MarginUsed            string `json:"marginUsed"`
}

type Account struct {
	ID      string         `json:"id"`
	Balance string         `json:"balance"`
	Orders  []AccountOrder `json:"orders"`
	Trades  []AccountTrade `json:"trades"`
}

type AccountDetail struct {
	Account Account `json:"account"`
}
