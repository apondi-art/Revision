package main

import "fmt"

func main() {
	size := 5
	fmt.Println(TwoDDimension(size))
}
func TwoDDimension(size int) [][]int {
	table := make([][]int, size)
	for i := range table {
		table[i] = make([]int, size)
	}
	for i := 1; i <= size; i++ {
		for j := 1; j <= size; j++ {
			table[i-1][j-1] = i * j
		}
	}
	return table
}
