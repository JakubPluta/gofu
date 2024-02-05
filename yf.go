package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const yfinanceQuotesUrl = `https://query2.finance.yahoo.com/v8/finance/chart/%s`
const qParams = "" //`?range=%s&interval=%s`

type QueryParams struct {
	Symbol   string
	Interval Interval
	Period   Period
}

func (y QueryParams) String() string {
	return fmt.Sprintf("%s,%s,%s", y.Symbol, y.Interval, y.Period)
}

// Params

type Interval string
type Period string

const (
	TF1m  Interval = "1m"
	TF2m  Interval = "2m"
	TF5m  Interval = "5m"
	TF15m Interval = "15m"
	TF30m Interval = "30m"
	TF60m Interval = "60m"
	TF90m Interval = "90m"
	TF1h  Interval = "1h"
	TF1d  Interval = "1d"
	TF5d  Interval = "5d"
	TF1wk Interval = "1wk"
	TF1mo Interval = "1mo"
	TF3mo Interval = "3mo"
)

const (
	TD1d  Period = "1d"
	TD5d  Period = "5d"
	TD7d  Period = "7d"
	TD1mo Period = "1mo"
	TD3mo Period = "3mo"
	TD6mo Period = "6mo"
	TD1y  Period = "1y"
	TD2y  Period = "2y"
	TD5y  Period = "5y"
	TD10y Period = "10y"
	TDytd Period = "ytd"
	TDmax Period = "max"
)

type YFResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Symbol string `json:"symbol"`
			}
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close  []float64 `json:"close"`
					High   []float64 `json:"high"`
					Volume []int     `json:"volume"`
					Open   []float64 `json:"open"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

type fullYFResponse struct {
	Chart struct {
		Result []struct {
			Meta struct {
				Currency             string  `json:"currency"`
				Symbol               string  `json:"symbol"`
				ExchangeName         string  `json:"exchangeName"`
				InstrumentType       string  `json:"instrumentType"`
				FirstTradeDate       int     `json:"firstTradeDate"`
				RegularMarketTime    int     `json:"regularMarketTime"`
				Gmtoffset            int     `json:"gmtoffset"`
				Timezone             string  `json:"timezone"`
				ExchangeTimezoneName string  `json:"exchangeTimezoneName"`
				RegularMarketPrice   float64 `json:"regularMarketPrice"`
				ChartPreviousClose   float64 `json:"chartPreviousClose"`
				PreviousClose        float64 `json:"previousClose"`
				Scale                int     `json:"scale"`
				PriceHint            int     `json:"priceHint"`
				CurrentTradingPeriod struct {
					Pre struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"pre"`
					Regular struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"regular"`
					Post struct {
						Timezone  string `json:"timezone"`
						End       int    `json:"end"`
						Start     int    `json:"start"`
						Gmtoffset int    `json:"gmtoffset"`
					} `json:"post"`
				} `json:"currentTradingPeriod"`
				TradingPeriods [][][]struct {
					Timezone  string `json:"timezone"`
					End       int    `json:"end"`
					Start     int    `json:"start"`
					Gmtoffset int    `json:"gmtoffset"`
				} `json:"tradingPeriods"`
				DataGranularity string   `json:"dataGranularity"`
				Range           string   `json:"range"`
				ValidRanges     []string `json:"validRanges"`
			} `json:"meta"`
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close  []float64 `json:"close"`
					High   []float64 `json:"high"`
					Volume []int     `json:"volume"`
					Open   []float64 `json:"open"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

// GetYahooFinanceQuotesData retrieves Yahoo Finance quotes data.
//
// symbol string, interval Interval, period Period.
// YFResponse.
func GetYahooFinanceQuotesData(symbol string, interval Interval, period Period) YFResponse {
	var response YFResponse

	client := http.Client{
		Timeout: 10 * time.Second,
	}
	//qp := fmt.Sprintf(qParams, interval, period)
	resp, err := client.Get(fmt.Sprintf(yfinanceQuotesUrl, strings.ToUpper(symbol)))
	if err != nil {
		log.Printf("Error: %s", err)
		return response
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Printf("Error: status code %d", resp.StatusCode)
		return response
	}

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error: %s", err)
		return response
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Printf("Error: %s", err)
		return response
	}

	return response
}
