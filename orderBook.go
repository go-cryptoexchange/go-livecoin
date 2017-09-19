package livecoin

type OrderBook struct {
	Timestamp int        `json:"timestamp"`
	Asks      [][]string `json:"asks"`
	Bids      [][]string `json:"bids"`
}
