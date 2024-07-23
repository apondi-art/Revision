package main

import "fmt"

func main() {
	s := [3]float64{1, 1, 1}
	v := 5
	fmt.Println(Tribonnaci(s, v))
}

// tribbonnacii means triple(adding 3 values to get the next values).
// fibonnanci means (adding to values to get the next value)
func Tribonnaci(s [3]float64, n int) []float64 {
	if n == 0 {
		return []float64{}
	}
	if n < 3 {
		return s[:n]
	}
	value := make([]float64, n)
	copy(value, s[:])
	for i := 3; i < len(value); i++ {
		value[i] = value[i-1] + value[i-2] + value[i-3]
	}
	return value
}
