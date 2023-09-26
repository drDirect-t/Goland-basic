package main

import "fmt"

func hello() {
	fmt.Println("Hello")
}

func plus(value1 int, value2 int) int {
	return value1 + value2

}

func value3(value1 int, value2 int, value3 int) int {
	return value1 + value2 + value3
}

func main() {
	//function 1
	hello()
	//function 2
	result := plus(20, 50)
	fmt.Println(result)
	//function 3
	result2 := value3(20, 50, 60)
	fmt.Println(result2)
}
