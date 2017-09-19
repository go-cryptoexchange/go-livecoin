package livecoin

import (
	"testing"
)

func TestGetTickers(t *testing.T) {
	api := New("", "")
	tickers, err := api.GetTickers()
	if err != nil {
		t.Fatal(err)
	}
	if len(tickers) < 1 {
		t.Fail()
	}
}

func TestGetOrderBook(t *testing.T) {
	api := New("", "")
	book, err := api.GetOrderBook("BTC/USD", false, 10)
	if err != nil {
		t.Fatal(err)
	}
	if len(book.Asks) < 1 {
		t.Fail()
	}
	if len(book.Bids) < 1 {
		t.Fail()
	}
}
