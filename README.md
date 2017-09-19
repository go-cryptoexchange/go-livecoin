# go-livecoin
Go wrapper around Livecoin API  

At that time only two public methods are implemented: `/exchange/ticker` and `/exchange/order_book`

## Usage

```go
package main

import (
    "log"
    "fmt"

    "github.com/go-cryptoexchange/go-livecoin"
)

func main() {
    api := livecoin.New("api-key", "api-secret")
    tickers, _ := api.GetTickers()
    for _, ticker := range tickers {
        log.Println(
            fmt.Sprintf("%+v", ticker),
        )
    }
}
```

Will output something like this:
```
2017/09/19 11:16:54 {Currency:MGO Symbol:MGO/BTC Last:0.00012407 High:0.00012407 Low:0.00012407 Volume:5.82341982 Vwap:0.00012407 Max_bid:0.00012469 Min_ask:0.00012406 Best_bid:0.00012469 Best_ask:0.000125}
2017/09/19 11:16:54 {Currency:MGO Symbol:MGO/ETH Last:0.00249 High:0.00249 Low:0.00249 Volume:3 Vwap:0.00249 Max_bid:0.00249 Min_ask:0 Best_bid:0.000249 Best_ask:0.00249}
2017/09/19 11:16:54 {Currency:CTR Symbol:CTR/USD Last:1.4 High:2.2 Low:1.4 Volume:166.54705534 Vwap:1.60510304 Max_bid:5.6 Min_ask:1.4 Best_bid:1.4 Best_ask:2.105}
```
