package main

type macd struct {
	SignalLine   []float64
	Line         []float64
	Histogram    []float64
	ShortPeriod  int
	LongPeriod   int
	SignalPeriod int
}

type sma struct {
	Value  []float64
	Window int
}

type ema struct {
	Value []float64
	Alpha float64
}

func sumValues[T int | float64](values []T) T {
	var sum T
	for _, v := range values {
		sum += v
	}
	return sum
}

func isValidAlpha(alpha float64) bool {
	return 0 <= alpha && alpha <= 1
}

func SMA(values []float64, window int) []float64 {
	var sma []float64
	size := len(values)
	start := size - window

	for i := range values[start:] {
		windowMean := sumValues(values[i:i+window]) / float64(window)
		sma = append(sma, windowMean)

	}
	return sma
}

func EMA(values []float64, alpha float64) []float64 {
	if !isValidAlpha(alpha) {
		panic("alpha must be between 0 and 1")
	}
	var ema []float64
	if len(values) == 0 {
		return ema
	}
	prevEMA := values[0]
	ema = append(ema, prevEMA)
	for i := 1; i < len(values); i++ {
		currentEMA := alpha*values[i] + (1-alpha)*prevEMA
		ema = append(ema, currentEMA)
		prevEMA = currentEMA
	}
	return ema
}

func MACD(values []float64, shortPeriod, longPeriod, signalPeriod int) ([]float64, []float64, []float64) {

	shortAdjustedAlpha := float64(2.0 / float64(shortPeriod+1))
	longAdjustedAlpha := float64(2.0 / float64(longPeriod+1))

	shortEMA := EMA(values, shortAdjustedAlpha)
	longEMA := EMA(values, longAdjustedAlpha)

	var macdLine []float64
	var macdHistogram []float64
	for i := range values {
		macdLine = append(macdLine, shortEMA[i]-longEMA[i])
	}
	signalLine := EMA(macdLine, float64(2.0/float64(signalPeriod+1)))

	for i := range values {
		macdHistogram = append(macdHistogram, macdLine[i]-signalLine[i])
	}
	return macdLine, macdHistogram, signalLine
}
