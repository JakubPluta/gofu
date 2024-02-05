package main

import "fmt"

func main() {
	// data := ReadCSV("data/example.csv")
	// quotes := ParseQuoteFromCSV(data)
	// fmt.Println(quotes)

	data := GetYahooFinanceQuotesData("META", "1m", "1y")
	_ = data
	fmt.Println(data)
	//fmt.Println(b)

}
