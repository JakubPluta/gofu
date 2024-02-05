package main

import (
	"fmt"
	"strconv"
	"time"
)

type OHLC struct {
	Date   time.Time
	Open   float64
	High   float64
	Low    float64
	Close  float64
	Volume float64
}

type Quotes []OHLC

func (o OHLC) String() string {
	return fmt.Sprintf("%s,%f,%f,%f,%f,%f", o.Date, o.Open, o.High, o.Low, o.Close, o.Volume)
}

func (q Quotes) String() string {
	s := ""
	s += "Date,Open,High,Low,Close,Volume\n"
	for _, ohlc := range q {
		s += ohlc.String() + "\n"
	}
	return s
}
func ParseQuotes(data [][]string) Quotes {
	quotes := Quotes{}
	for _, line := range data[1:] {
		ohlc := OHLC{}
		ohlc.Date, _ = time.Parse("2006-01-02", line[0])
		ohlc.Open, _ = strconv.ParseFloat(line[1], 64)
		ohlc.High, _ = strconv.ParseFloat(line[2], 64)
		ohlc.Low, _ = strconv.ParseFloat(line[3], 64)
		ohlc.Close, _ = strconv.ParseFloat(line[4], 64)
		ohlc.Volume, _ = strconv.ParseFloat(line[5], 64)
		quotes = append(quotes, ohlc)
	}
	return quotes
}

func ParseQuotesFromJSON(data map[string]interface{}) Quotes {

}
