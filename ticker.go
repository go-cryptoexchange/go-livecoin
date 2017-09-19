package livecoin

type Ticker struct {
	Currency string  `json:"cur"`
	Symbol   string  `json:"symbol"`
	Last     float64 `json:"last"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Volume   float64 `json:"volume"`
	Vwap     float64 `json:"vwap"`
	Max_bid  float64 `json:"max_bid"`
	Min_ask  float64 `json:"min_ask"`
	Best_bid float64 `json:"best_bid"`
	Best_ask float64 `json:"best_ask"`
}
