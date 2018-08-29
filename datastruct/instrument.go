package datastruct

type Instrument struct {
	Name                        string  `json:"name"`
	Type                        string  `json:"type"`
	DisplayName                 string  `json:"displayName"`
	PipLocation                 float64 `json:"pipLocation"`
	DisplayPrecision            float64 `json:"displayPrecision"`
	TradeUnitsPrecision         float64 `json:"tradeUnitsPrecision"`
	MinimumTradeSize            string  `json:"minimumTradeSize"`
	MaximumTrailingStopDistance string  `json:"maximumTrailingStopDistance"`
	MinimumTrailingStopDistance string  `json:"minimumTrailingStopDistance"`
	MaximumPositionSize         string  `json:"maximumPositionSize"`
	MaximumOrderUnits           string  `json:"maximumOrderUnits"`
	MarginRate                  string  `json:"marginRate"`
}

type Instruments struct {
	Instruments []Instrument `json:"instruments"`
}

type Candle struct {
	Complete bool   `json:"complete"`
	Volume   int    `json:"volume"`
	Time     string `json:"time"`
	Mid      map[string]string
}

type Candles struct {
	Instrument  string `json:"instrument"`
	Granularity string `json:"granularity"`
	Candles     []Candle
}

type CandleQueryParam struct {
	Granularity string `json:"granularity"`
	Count int `json:"count"`
}