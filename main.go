package main

import "fmt"

func main() {
	// data := ReadCSV("data/example.csv")
	// quotes := ParseQuoteFromCSV(data)
	// fmt.Println(quotes)

	data := GetYahooFinanceQuotesData("META", TF1h, "1d")
	_ = data
	fmt.Println(data)
	//fmt.Println(b)

}
