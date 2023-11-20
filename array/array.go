package main

import (
	"fmt"
)

func main() {
	a := []int{5, 10, 165, 7, 0, 633, 156, 103}
	sum := 0
	for _, v := range a {
		sum += v
	}
	fmt.Println(sum)
}
