package main

import "fmt"

func pointer(ivalue int) {
	ivalue = 0
}

func ipointer(iPointer *int) {
	*iPointer = 0
}

func main() {
	i := 1
	fmt.Println("i =", i)

	pointer(i)
	fmt.Println("i =", i)

	ipointer(&i)
	fmt.Println("i value from ipointer =", i)
	fmt.Println("i address from ipointer =", &i)
}
