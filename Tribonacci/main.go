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
	value := make([]float64, n)
	copy(value, s[:])
	for i := 4; i < len(value); i++ {
		value[i] = value[i-1] + value[i-2] + value[i-3]
	}
	return value
}
