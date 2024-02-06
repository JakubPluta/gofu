package main

import "fmt"

func main() {
	// data := ReadCSV("data/example.csv")
	// quotes := ParseQuoteFromCSV(data)
	// fmt.Println(quotes)

	data := RetrieveDailyOHLCData("META", "30d")
	_ = data
	for _, ohlc := range data {
		fmt.Println(ohlc)
	}

}
