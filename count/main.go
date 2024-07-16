package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := []int{21, 30, 34, -30, -25, -90, 40}
	fmt.Println(CountResult(s))
}
func CountResult(s []int) string {
	var result string
	positive := 0
	negative := 0
	for _, el := range s {
		if el > 0 {
			positive++
		} else {
			negative++
		}
	}
	result += strconv.Itoa(positive)
	result += "," + strconv.Itoa(negative)
	return result
}
