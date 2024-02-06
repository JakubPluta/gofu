package main

import (
	"fmt"
	"time"
)

type Quotes struct {
	Date   []time.Time `json:"date"`
	Open   []float64   `json:"open"`
	High   []float64   `json:"high"`
	Low    []float64   `json:"low"`
	Close  []float64   `json:"close"`
	Volume []float64   `json:"volume"`
}

type OHLC struct {
	Date   time.Time `json:"date"`
	Open   float64   `json:"open"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Close  float64   `json:"close"`
	Volume float64   `json:"volume"`
}

type OHLCS []OHLC

func (o OHLC) String() string {
	return fmt.Sprintf("%s,%f,%f,%f,%f,%f", o.Date, o.Open, o.High, o.Low, o.Close, o.Volume)
}

func (q Quotes) ToOHLC() []OHLC {
	var ohlc []OHLC
	for i := range q.Date {
		ohlc = append(ohlc, OHLC{
			Date:   q.Date[i],
			Open:   q.Open[i],
			High:   q.High[i],
			Low:    q.Low[i],
			Close:  q.Close[i],
			Volume: q.Volume[i],
		})
	}
	return ohlc
}

func (o OHLCS) ToQuotes() Quotes {
	var quotes Quotes
	for i := range o {
		quotes.Date = append(quotes.Date, o[i].Date)
		quotes.Open = append(quotes.Open, o[i].Open)
		quotes.High = append(quotes.High, o[i].High)
		quotes.Low = append(quotes.Low, o[i].Low)
		quotes.Close = append(quotes.Close, o[i].Close)
		quotes.Volume = append(quotes.Volume, o[i].Volume)
	}
	return quotes
}

func (o OHLCS) String() string {
	s := ""
	s += "Date,Open,High,Low,Close,Volume\n"
	for _, ohlc := range o {
		s += ohlc.String() + "\n"
	}
	return s
}

func (o OHLCS) Len() int {
	return len(o)
}

func (o Quotes) Len() int {
	return len(o.Date)
}

func (o Quotes) String() string {
	s := ""
	s += "Date,Open,High,Low,Close,Volume\n"
	for _, ohlc := range o.ToOHLC() {
		s += ohlc.String() + "\n"
	}
	return s
}
