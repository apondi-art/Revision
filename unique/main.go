package main

import "fmt"

func main() {
	s := [5]int{2, 4, 6, 2, 4}
	fmt.Println(Unique(s))
}
func Unique(s [5]int) int {
	mapv := make(map[int]int)
	for _, number := range s {
		mapv[number]++
	}
	for k, v := range mapv {
		if v == 1 {
			return k
		}
	}
	return 0
}
