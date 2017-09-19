package livecoin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const (
	API_BASE = "https://api.livecoin.net"
)

// livecoin represent a livecoin client
type Livecoin struct {
	client *client
}

// New returns an instantiated livecoin struct
func New(apiKey, apiSecret string) *Livecoin {
	client := NewClient(apiKey, apiSecret)
	return &Livecoin{client}
}

// NewWithCustomHttpClient returns an instantiated livecoin struct with custom http client
func NewWithCustomHttpClient(apiKey, apiSecret string, httpClient *http.Client) *Livecoin {
	client := NewClientWithCustomHttpConfig(apiKey, apiSecret, httpClient)
	return &Livecoin{client}
}

// NewWithCustomTimeout returns an instantiated livecoin struct with custom timeout
func NewWithCustomTimeout(apiKey, apiSecret string, timeout time.Duration) *Livecoin {
	client := NewClientWithCustomTimeout(apiKey, apiSecret, timeout)
	return &Livecoin{client}
}

func (l *Livecoin) GetTickers() ([]Ticker, error) {
	var tickers []Ticker
	r, err := l.client.do("GET", "/exchange/ticker", "")
	if err != nil {
		return tickers, err
	}
	err = json.Unmarshal(r, &tickers)
	if err != nil {
		return tickers, err
	}
	return tickers, nil
}

func (l *Livecoin) GetOrderBook(currencyPair string, groupByPrice bool, depth int) (OrderBook, error) {
	var orderBook OrderBook
	r, err := l.client.do(
		"GET",
		fmt.Sprintf("/exchange/order_book?currencyPair=%s&groupByPrice=%s&depth=%s", currencyPair, groupByPrice, depth),
		"",
	)
	if err != nil {
		return orderBook, err
	}
	err = json.Unmarshal(r, &orderBook)
	if err != nil {
		return orderBook, err
	}
	return orderBook, nil
}
