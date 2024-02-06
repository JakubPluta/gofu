package main

import "fmt"

func main() {

	data := RetrieveDailyOHLCData("META", "30d")
	_ = data
	for _, ohlc := range data {
		fmt.Println(ohlc)
	}

	s := EMA(data.ToQuotes().Close, 0.3)
	fmt.Println(s)

	m, s, t := MACD(data.ToQuotes().Close, 5, 8, 9)
	fmt.Println(m, s, t)
}
