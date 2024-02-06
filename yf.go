package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const yfinanceQuotesUrl = `https://query2.finance.yahoo.com/v8/finance/chart/%s?interval=%s&range=%s`

func isValidInterval[T comparable](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

var availableIntervals = []string{"1m", "2m", "5m", "15m", "30m", "60m", "90m", "1h", "1d", "5d", "1wk", "1mo", "3mo"}

func convertMonthsToDays(period string) (int, error) {
	months := strings.Split(period, "mo")
	numMonths, err := strconv.Atoi(months[0])
	if err != nil {
		return 0, err
	}
	return numMonths * 30, nil
}

func convertYearsToDays(period string) (int, error) {
	years := strings.Split(period, "y")
	numYears, err := strconv.Atoi(years[0])
	if err != nil {
		return 0, err
	}
	return numYears * 365, nil
}

type ErrorResponse struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}
type YFResponse struct {
	Chart struct {
		Result []struct {
			ErrorResponse `json:"error"`
			Meta          struct {
				Symbol string `json:"symbol"`
			}
			Timestamp  []int `json:"timestamp"`
			Indicators struct {
				Quote []struct {
					Close  []float64 `json:"close"`
					High   []float64 `json:"high"`
					Volume []int     `json:"volume"`
					Open   []float64 `json:"open"`
					Low    []float64 `json:"low"`
				} `json:"quote"`
			} `json:"indicators"`
		} `json:"result"`
	} `json:"chart"`
}

// symbol: string - the symbol for the stock or asset
// period: string - the time period for the data
// []OHLC - an array of OHLC (open, high, low, close) data points
func RetrieveDailyOHLCData(symbol string, period string) []OHLC {
	resp, err := getYahooFinanceQuotesData(symbol, period, "1d")
	if err != nil {
		log.Fatal(err)
	}
	quotes := resp.Chart.Result[0].Indicators.Quote[0]
	timestamps := resp.Chart.Result[0].Timestamp

	var dailyOHLCData []OHLC

	for i, ts := range timestamps {
		dailyOHLCData = append(dailyOHLCData, OHLC{
			Date:   time.Unix(int64(ts), 0),
			Open:   quotes.Open[i],
			High:   quotes.High[i],
			Low:    quotes.Low[i],
			Close:  quotes.Close[i],
			Volume: float64(quotes.Volume[i]),
		})
	}

	return dailyOHLCData
}

func getYahooFinanceQuotesData(symbol string, period string, interval string) (YFResponse, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	url := fmt.Sprintf(yfinanceQuotesUrl, strings.ToUpper(symbol), interval, period)
	resp, err := client.Get(url)
	if err != nil {
		return YFResponse{}, fmt.Errorf("Error retrieving Yahoo Finance quotes data: %s", err)
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		return YFResponse{}, fmt.Errorf("Error reading response data: %s", err)
	}
	if resp.StatusCode != http.StatusOK {
		return YFResponse{}, fmt.Errorf("Error: status code %d - %s, body: %s", resp.StatusCode, resp.Status, string(responseData))
	}

	var response YFResponse
	err = json.Unmarshal(responseData, &response)
	if err != nil {
		return YFResponse{}, fmt.Errorf("Error unmarshalling response data: %s", err)
	}

	return response, nil
}
