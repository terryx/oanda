package datastruct

type StopLossOnFill struct {
	Price string `json:"price"`
}

type TakeProfitOnFill struct {
	Price string `json:"price"`
}

type OrderForm struct {
	Type             string `json:"type"`
	Instrument       string `json:"instrument"`
	Units            int    `json:"units"`
	Price            string `json:"price"`
	TimeInForce      string `json:"timeInForce"`
	PositionFill     string `json:"positionFill"`
	StopLossOnFill   `json:"stopLossOnFill"`
	TakeProfitOnFill `json:"takeProfitOnFill"`
}

type OrderCreateResponse struct {
	Transaction struct {
		Type             string `json:"type"`
		Instrument       string `json:"instrument"`
		Units            string `json:"units"`
		Price            string `json:"price"`
		TimeInForce      string `json:"timeInForce"`
		TriggerCondition string `json:"triggerCondition"`
		PartialFill      string `json:"partialFill"`
		PositionFill     string `json:"positionFill"`
		StopLossOnFill   struct {
			Price       string `json:"price"`
			TimeInForce string `json:"timeInForce"`
		} `json:"stopLossOnFill"`
		TrailingStopLossOnFill struct {
			Distance    string `json:"distance"`
			TimeInForce string `json:"timeInForce"`
		} `json:"trailingStopLossOnFill"`
		Reason    string `json:"reason"`
		ID        string `json:"id"`
		UserID    int    `json:"UserID"`
		AccountID string `json:"accountID"`
		BatchID   string `json:"batchID"`
		RequestID string `json:"requestID"`
		Time      string `json:"time"`
	} `json:"orderCreateTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

type OrderCancelResponse struct {
	Transaction struct {
		Type      string `json:"type"`
		OrderID   string `json:"orderID"`
		Reason    string `json:"reason"`
		ID        string `json:"id"`
		UserID    int    `json:"UserID"`
		AccountID string `json:"accountID"`
		BatchID   string `json:"batchID"`
		RequestID string `json:"requestID"`
		Time      string `json:"time"`
	} `json:"orderCancelTransaction"`
	RelatedTransactionIDs []string `json:"relatedTransactionIDs"`
	LastTransactionID     string   `json:"lastTransactionID"`
}

