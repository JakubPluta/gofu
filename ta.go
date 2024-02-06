package main

type Indicator struct {
	Name   string
	Func   func([]float64) float64
	Params []float64
}

func SMA(data []float64, period int) float64 {
	sum := 0.0
	for i := 0; i < period; i++ {
		sum += data[i]
	}
	return sum / float64(period)
}
