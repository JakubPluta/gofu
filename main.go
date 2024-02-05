package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	data := ReadCSV("yahoofinance-SPY-20200901-20210113.csv")
	quotes := ParseQuotes(data)
	fmt.Println(quotes)
}
