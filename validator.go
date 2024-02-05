package main

func validateOHLC(ohlc []float64) bool {
	if len(ohlc) != 3 {
		return false
	}
	return true
}
