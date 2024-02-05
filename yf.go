package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const yfinanceQuotesUrl = `https://query2.finance.yahoo.com/v8/finance/chart/%s?interval=%s&range=%s`

// Params

type Interval string
type Period string

func (i Interval) String() string {
	return string(i)
}

func (p Period) String() string {
	return string(p)
}

// Contains checks if the given element is present in the array of Intervals.
//
// arr []Interval - the array of Intervals to search
// str Interval - the Interval to search for
// bool - true if the Interval is found, false otherwise
func isValidInterval[T comparable](arr []T, str T) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func (p Period) toDays() (string, error) {
	period := p.String()
	if strings.HasSuffix(period, "d") {
		return period, nil
	}
	if strings.HasSuffix(period, "m") || strings.HasSuffix(period, "h") || p == "max" {
		return "", errors.New("cannot convert minutes/hours to days")
	}
	if strings.HasSuffix(period, "mo") {
		// convert string to slice
		months := strings.Split(period, "mo")
		numMonths, err := strconv.Atoi(months[0])
		if err != nil {
			return "", err
		}
		return strconv.Itoa(numMonths * 30), nil

	}
	if strings.HasSuffix(period, "y") {
		// convert string to slice
		years := strings.Split(period, "y")
		numYears, err := strconv.Atoi(years[0])
		if err != nil {
			return "", err
		}
		return strconv.Itoa(numYears * 365), nil
	}

	return period, nil
}

func isValidPeriodForInterval(interval, period string) bool {
	if period == "max" {
		return true
	}

	p := Period(period)
	days, err := p.toDays()
	if err != nil {
		log.Fatal(err)
	}
	if days == "" {
		return false
	}

	daysSlice := strings.Split(days, "d")
	dayNum, err := strconv.Atoi(daysSlice[0])
	if err != nil {
		log.Fatal(err)
	}
	if interval == "1m" && dayNum > 7 {
		return false
	}
	if (interval == "2m" || interval == "5m" || interval == "15m" || interval == "30m" || interval == "90m") && dayNum > 60 {
		return false
	}

	if (interval == "60m" || interval == "1h") && dayNum > 730 {
		return false
	}

	return true
}

type YFResponse struct {
	Chart struct {
		Result []struct {
			Error struct {
				Code        string `json:"code`
				Description string `json:"description"`
			} `json:"error"`
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

type FullYFResponse struct {
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

func GetYahooFinanceQuotesData(symbol string, interval Interval, period Period) YFResponse {
	var availableIntervals = []Interval{"1m", "2m", "5m", "15m", "30m", "60m", "90m", "1h", "1d", "5d", "1wk", "1mo", "3mo"}
	var response YFResponse

	if !isValidInterval(availableIntervals, interval) {
		log.Printf("Invalid interval: %s. Available intervals: %v. Setting default interval: %s", interval, availableIntervals, "1d")
		interval = "1d"
	}

	client := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := client.Get(fmt.Sprintf(yfinanceQuotesUrl, strings.ToUpper(symbol), interval, period))
	if err != nil {
		log.Printf("Error retrieving Yahoo Finance quotes data: %s", err)
		return response
	}
	defer resp.Body.Close()

	responseData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response data: %s", err)
		return response
	}
	if resp.StatusCode != http.StatusOK {
		log.Printf("Error: status code %d - %s, body: %s", resp.StatusCode, resp.Status, string(responseData))
		return response
	}

	err = json.Unmarshal(responseData, &response)
	if err != nil {
		log.Printf("Error unmarshalling response data: %s", err)
		return response
	}

	return response
}
